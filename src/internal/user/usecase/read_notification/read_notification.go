package read_notification

import (
	"context"
)

type Repository interface {
	MarkReadNotification(userId int64, notificationId int64) error
}

type UseCaseReadNotification struct {
	repository Repository
}

func NewUseCaseReadNotification(repository Repository) UseCaseReadNotification {
	return UseCaseReadNotification{repository: repository}
}

func (uc UseCaseReadNotification) Execute(ctx context.Context, userId int64, notificationId int64) error {
	err := uc.repository.MarkReadNotification(userId, notificationId)
	if err != nil {
		return err
	}
	return nil
}
