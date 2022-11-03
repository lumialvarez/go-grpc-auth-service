package validate

import (
	"context"
	domainError "github.com/lumialvarez/go-grpc-auth-service/src/internal/error"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user"
)

type Repository interface {
	GetByEmail(email string) (*user.User, error)
	Save(user *user.User) error
}

type JwtServiceUser interface {
	GenerateToken(user *user.User) (signedToken string, err error)
	ValidateToken(signedToken string) (*user.User, error)
}

type UseCaseValidateUser struct {
	repository Repository
	jwtService JwtServiceUser
}

func NewUseCaseValidateUser(repository Repository, jwtService JwtServiceUser) UseCaseValidateUser {
	return UseCaseValidateUser{repository: repository, jwtService: jwtService}
}

func (uc UseCaseValidateUser) Execute(ctx context.Context, domainUser *user.User) (*user.User, error) {
	jwtUser, err := uc.jwtService.ValidateToken(domainUser.Token())

	if err != nil {
		return nil, domainError.NewInvalidCredentials("Invalid Token")
	}

	dbUser, err := uc.repository.GetByEmail(jwtUser.Email())
	if err != nil {
		return nil, domainError.NewInvalidCredentials("Invalid Token")
	}

	if dbUser.Id() != jwtUser.Id() {
		return nil, domainError.NewInvalidCredentials("Invalid Token")
	}

	return dbUser, nil
}
