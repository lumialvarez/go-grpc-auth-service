package register

import (
	"context"
	"github.com/lumialvarez/go-common-tools/hash"
	domainError "github.com/lumialvarez/go-grpc-auth-service/src/internal/error"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user"
)

type Repository interface {
	GetByEmail(email string) (*user.User, error)
	GetByUserName(username string) (*user.User, error)
	Save(user *user.User) (*user.User, error)
}

type UseCaseRegisterUser struct {
	repository Repository
}

func NewUseCaseRegisterUser(repository Repository) UseCaseRegisterUser {
	return UseCaseRegisterUser{repository: repository}
}

func (uc UseCaseRegisterUser) Execute(ctx context.Context, domainUser *user.User) (*user.User, error) {
	domainUser.SetPassword(hash.HashPassword(domainUser.Password()))
	_, err := uc.repository.GetByUserName(domainUser.UserName())
	if err == nil {
		return nil, domainError.NewAlreadyExists("User Name already exists")
	}

	createdUser, err := uc.repository.Save(domainUser)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
