package mapper

import (
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/repository/postgresql/user/dao"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user"
)

type Mapper struct {
}

func (m Mapper) ToDomain(daoUser *dao.User) *user.User {
	return user.NewUser(daoUser.Id, daoUser.Email, daoUser.Password)
}

func (m Mapper) ToDAO(domainUser *user.User) *dao.User {
	daoUser := dao.User{
		Id:       domainUser.Id(),
		Email:    domainUser.Email(),
		Password: domainUser.Password(),
	}
	return &daoUser
}
