package mapper

import (
	"github.com/lumialvarez/go-common-tools/service/rabbitmq/notification/dto"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user"
	"time"
)

type Mapper struct {
}

func (m Mapper) ToDomain(dtoNotification dto.Notification) *user.Notification {

	domainNotification := user.NewNotification(0, dtoNotification.Title, dtoNotification.Detail, time.Now(), false)

	return domainNotification
}
