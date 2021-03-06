package login

import (
	"context"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/utils"
	domainError "github.com/lumialvarez/go-grpc-auth-service/src/internal/error"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user"
)

type Repository interface {
	GetByEmail(email string) (*user.User, error)
	Save(user *user.User) error
}

type JwtServiceUser interface {
	GenerateToken(user *user.User) (signedToken string, err error)
	ValidateToken(signedToken string) (err error)
}

type UseCaseLoginUser struct {
	repository Repository
	jwtService JwtServiceUser
}

func NewUseCaseLoginUser(repository Repository, jwtService JwtServiceUser) UseCaseLoginUser {
	return UseCaseLoginUser{repository: repository, jwtService: jwtService}
}

func (uc UseCaseLoginUser) Execute(ctx context.Context, domainUser *user.User) (*user.User, error) {
	dbUser, err := uc.repository.GetByEmail(domainUser.Email())
	if err != nil {
		return nil, domainError.NewNotFound("User not found")
	}

	match := utils.CheckPasswordHash(domainUser.Password(), dbUser.Password())

	if !match {
		return nil, domainError.NewNotFound("User not found")
	}

	token, _ := uc.jwtService.GenerateToken(dbUser)
	domainUser.SetToken(token)

	return domainUser, nil
}
