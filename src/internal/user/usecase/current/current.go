package current

import (
	"context"
	domainError "github.com/lumialvarez/go-grpc-auth-service/src/internal/error"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user"
)

type Repository interface {
	GetById(id int64) (*user.User, error)
}

type JwtServiceUser interface {
	ValidateToken(signedToken string) (*user.User, error)
}

type UseCaseCurrentUser struct {
	repository Repository
	jwtService JwtServiceUser
}

func NewUseCaseCurrentUser(repository Repository, jwtService JwtServiceUser) UseCaseCurrentUser {
	return UseCaseCurrentUser{repository: repository, jwtService: jwtService}
}

func (uc UseCaseCurrentUser) Execute(ctx context.Context, token string) (*user.User, error) {
	jwtUser, err := uc.jwtService.ValidateToken(token)
	if err != nil {
		return nil, domainError.NewInvalidCredentials("Invalid Token")
	}

	dbUser, err := uc.repository.GetById(jwtUser.Id())
	if err != nil {
		return nil, domainError.NewInvalidCredentials("Invalid Token")
	}

	return dbUser, nil
}
