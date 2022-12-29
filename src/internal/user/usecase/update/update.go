package update

import (
	"context"
	"github.com/lumialvarez/go-common-tools/hash"
	domainError "github.com/lumialvarez/go-grpc-auth-service/src/internal/error"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user"
)

type Repository interface {
	GetByUserName(username string) (*user.User, error)
	Update(domainUser *user.User) (*user.User, error)
}

type UseCaseUpdateUser struct {
	repository Repository
}

func NewUseCaseUpdateUser(repository Repository) UseCaseUpdateUser {
	return UseCaseUpdateUser{repository: repository}
}

func (uc UseCaseUpdateUser) Execute(ctx context.Context, domainUser *user.User) (*user.User, error) {
	dbUser, err := uc.repository.GetByUserName(domainUser.UserName())
	if err != nil {
		return nil, domainError.NewNotFound("Username no exists")
	}
	if domainUser.Id() != dbUser.Id() {
		return nil, domainError.NewNotFound("Invalid User ID")
	}

	if len(domainUser.Password()) > 0 {
		domainUser.SetPassword(hash.HashPassword(domainUser.Password()))
	}

	createdUser, err := uc.repository.Update(domainUser)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
