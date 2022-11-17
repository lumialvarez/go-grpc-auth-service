package login

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

type JwtServiceUser interface {
	GenerateToken(user *user.User) (signedToken string, err error)
	ValidateToken(signedToken string) (*user.User, error)
}

type UseCaseLoginUser struct {
	repository Repository
	jwtService JwtServiceUser
}

func NewUseCaseLoginUser(repository Repository, jwtService JwtServiceUser) UseCaseLoginUser {
	return UseCaseLoginUser{repository: repository, jwtService: jwtService}
}

func (uc UseCaseLoginUser) Execute(ctx context.Context, domainUser *user.User) (*user.User, error) {
	dbUser, err := uc.repository.GetByUserName(domainUser.UserName())
	if err != nil {
		return nil, domainError.NewInvalidCredentials("Invalid credentials")
	}

	match := hash.CheckPasswordHash(domainUser.Password(), dbUser.Password())

	if !match {
		return nil, domainError.NewInvalidCredentials("Invalid credentials")
	}

	token, _ := uc.jwtService.GenerateToken(dbUser)
	dbUser.SetToken(token)

	return dbUser, nil
}
