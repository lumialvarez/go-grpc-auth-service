package notification

import (
	"encoding/json"
	"github.com/lumialvarez/go-common-tools/platform/rabbitmq"
	"github.com/lumialvarez/go-common-tools/service/rabbitmq/notification/dto"
	"github.com/lumialvarez/go-grpc-auth-service/src/cmd/devapi/config"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/consumers/notification/mapper"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type UseCaseCreateNotification interface {
	Execute(notification user.Notification) error
}

type Consumer struct {
	useCaseCreateNotification UseCaseCreateNotification
	mapper.Mapper
}

func NewConsumer(useCaseCreateNotification UseCaseCreateNotification) Consumer {
	return Consumer{useCaseCreateNotification: useCaseCreateNotification}
}

func (consumer Consumer) Init(config config.Config) {
	rabbitmqClient := rabbitmq.Init(config.RabbitMQUrl)

	q, err := rabbitmqClient.Channel.QueueDeclare(
		"NOTIFICATION_QUEUE", // name
		true,                 // durable
		false,                // delete when unused
		false,                // exclusive
		false,                // no-wait
		nil,                  // arguments
	)
	if err != nil {
		log.Print(err, "Failed to declare a queue Consumer")
	}

	msgs, err := rabbitmqClient.Channel.Consume(
		q.Name,                  // queue
		"Notification Consumer", // consumer
		false,                   // auto-ack
		false,                   // exclusive
		false,                   // no-local
		false,                   // no-wait
		nil,                     // args
	)
	if err != nil {
		log.Print(err, "Failed to register a consumer")
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {

			err, requeue := consumer.handleNotificationMessage(d)
			if err != nil {
				log.Print("Failed to handle notification message", err)
				d.Nack(false, requeue)

				time.Sleep(500 * time.Millisecond)
				continue
			}
			d.Ack(false)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	log.Printf(" [*] Waiting for notifications messages.")
	<-forever
}

func (consumer Consumer) handleNotificationMessage(delivery amqp.Delivery) (error, bool) {
	request := dto.Notification{}
	err := json.Unmarshal(delivery.Body, &request)
	if err != nil {
		return err, false
	}

	domainNotification := consumer.ToDomain(request)

	err = consumer.useCaseCreateNotification.Execute(*domainNotification)
	if err != nil {
		return err, true
	}

	return nil, true
}
