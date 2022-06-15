package user

import (
	"github.com/lumialvarez/go-grpc-auth-service/src/cmd/devapi/config"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/dto"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/platform/postgresql"
)

type Repository struct {
	postgresql postgresql.Client
	//mapper Mapper
}

func Init(config config.Config) Repository {
	return Repository{postgresql: postgresql.Init(config.DBUrl)}
}

func (repository *Repository) GetByEmail(email string) (*dto.User, error) {
	var user dto.User
	result := repository.postgresql.DB.Where(&dto.User{Email: email}).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repository *Repository) Save(user *dto.User) error {
	result := repository.postgresql.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
