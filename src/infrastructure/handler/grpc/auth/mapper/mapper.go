package mapper

import (
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/pb"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/repository/postgresql/user/dao"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user"
)

type Mapper struct {
}

func (m Mapper) ToDomainLogin(loginReq *pb.LoginRequest) *user.User {
	return user.NewUser(0, loginReq.Email, loginReq.Password)
}

func (m Mapper) ToDTO(domainUser *user.User) *dao.User {
	daoUser := dao.User{
		Id:       domainUser.Id(),
		Email:    domainUser.Email(),
		Password: domainUser.Password(),
	}
	return &daoUser
}
