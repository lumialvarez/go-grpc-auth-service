package notification

import (
	"fmt"
	"github.com/lumialvarez/go-common-tools/platform/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"time"
)

type Consumer struct {
}

func (consumer Consumer) Init() {
	rabbitUrl := os.Getenv("RABBITMQ_URL")
	rabbitmqClient := rabbitmq.Init(rabbitUrl)

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

			err := handleNotificationMessage(d)
			if err != nil {
				log.Print("Failed to handle notification message")
				continue
			}
			d.Ack(false)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func handleNotificationMessage(delivery amqp.Delivery) error {
	body := string(delivery.Body)

	fmt.Println(body)

	return nil
}
