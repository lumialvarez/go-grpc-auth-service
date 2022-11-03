package mapper

import (
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/pb"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user"
)

type Mapper struct {
}

func (m Mapper) ToDomainRegister(registerReq *pb.RegisterRequest) *user.User {
	return user.NewUser(0, registerReq.Email, registerReq.Password, "")
}

func (m Mapper) ToDomainLogin(loginReq *pb.LoginRequest) *user.User {
	return user.NewUser(0, loginReq.Email, loginReq.Password, "")
}

func (m Mapper) ToDomainValidate(validateReq *pb.ValidateRequest) *user.User {
	return user.NewUser(0, "", "", validateReq.Token)
}
