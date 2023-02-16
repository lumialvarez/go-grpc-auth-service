package create_notification

import (
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user"
)

type Repository interface {
	CreateNotificationToAdminUsers(notification user.Notification) error
}

type UseCaseCreateNotification struct {
	repository Repository
}

func NewUseCaseCreateNotification(repository Repository) UseCaseCreateNotification {
	return UseCaseCreateNotification{repository: repository}
}

func (uc UseCaseCreateNotification) Execute(notification user.Notification) error {
	err := uc.repository.CreateNotificationToAdminUsers(notification)
	if err != nil {
		return err
	}
	return nil
}
