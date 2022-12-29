package list

import (
	"context"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user"
)

type Repository interface {
	GetAll() (*[]user.User, error)
}

type UseCaseListUser struct {
	repository Repository
}

func NewUseCaseListUser(repository Repository) UseCaseListUser {
	return UseCaseListUser{repository: repository}
}

func (uc UseCaseListUser) Execute(ctx context.Context) (*[]user.User, error) {
	domainUsers, err := uc.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return domainUsers, nil
}
