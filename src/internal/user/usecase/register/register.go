package register

import (
	"context"
	domainError "github.com/lumialvarez/go-grpc-auth-service/src/internal/error"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user"
)

type Repository interface {
	GetByEmail(email string) (*user.User, error)
	Save(user *user.User) error
}

type UseCaseRegisterUser struct {
	repository Repository
}

func NewUseCaseRegisterUser(repository Repository) UseCaseRegisterUser {
	return UseCaseRegisterUser{repository: repository}
}

func (uc UseCaseRegisterUser) Execute(ctx context.Context, domainUser *user.User) error {
	user, err := uc.repository.GetByEmail(domainUser.Email())
	if err == nil {
		return domainError.NewAlreadyExists("E-Mail already exists")
	}

	err = uc.repository.Save(user)
	if err != nil {
		return err
	}

	return nil
}
