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
	Save(user *user.User) error
}

type JwtServiceUser interface {
	GenerateToken(user *user.User) (signedToken string, err error)
	ValidateToken(signedToken string) (*user.User, error)
}

type UseCaseRegisterUser struct {
	repository Repository
	jwtService JwtServiceUser
}

func NewUseCaseRegisterUser(repository Repository, jwtService JwtServiceUser) UseCaseRegisterUser {
	return UseCaseRegisterUser{repository: repository, jwtService: jwtService}
}

func (uc UseCaseRegisterUser) Execute(ctx context.Context, domainUser *user.User) error {
	domainUser.SetPassword(hash.HashPassword(domainUser.Password()))
	_, err := uc.repository.GetByUserName(domainUser.UserName())
	if err == nil {
		return domainError.NewAlreadyExists("User Name already exists")
	}

	err = uc.repository.Save(domainUser)
	if err != nil {
		return err
	}

	return nil
}
