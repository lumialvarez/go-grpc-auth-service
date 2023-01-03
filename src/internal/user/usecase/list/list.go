package list

import (
	"context"
	domainError "github.com/lumialvarez/go-grpc-auth-service/src/internal/error"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user"
)

type Repository interface {
	GetAll() (*[]user.User, error)
	GetById(id int64) (*user.User, error)
	GetByUserName(username string) (*user.User, error)
}

type UseCaseListUser struct {
	repository Repository
}

func NewUseCaseListUser(repository Repository) UseCaseListUser {
	return UseCaseListUser{repository: repository}
}

func (uc UseCaseListUser) Execute(ctx context.Context, id int64, userName string) (*[]user.User, error) {

	if id > 0 {
		var domainUsers []user.User
		domainUser, err := uc.repository.GetById(id)
		if err != nil {
			return nil, err
		}
		if len(userName) > 0 && domainUser.UserName() != userName {
			return nil, domainError.NewNotFound("User ID and Username mismatch")
		}
		domainUsers = append(domainUsers, *domainUser)
		return &domainUsers, nil
	}

	if len(userName) > 0 {
		var domainUsers []user.User
		domainUser, err := uc.repository.GetByUserName(userName)
		if err != nil {
			return nil, err
		}
		domainUsers = append(domainUsers, *domainUser)
		return &domainUsers, nil
	}

	domainUsers, err := uc.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return domainUsers, nil
}
