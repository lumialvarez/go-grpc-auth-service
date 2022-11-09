package mapper

import (
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/pb"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user"
)

type Mapper struct {
}

func (m Mapper) ToDomainRegister(registerReq *pb.RegisterRequest) *user.User {
	return user.NewUser(0, registerReq.Name, registerReq.UserName, registerReq.Email, registerReq.Password, "", m.ToDomainRole(registerReq.Role))
}

func (m Mapper) ToDomainLogin(loginReq *pb.LoginRequest) *user.User {
	return user.NewUser(0, "", loginReq.UserName, "", loginReq.Password, "", "")
}

func (m Mapper) ToDomainValidate(validateReq *pb.ValidateRequest) *user.User {
	return user.NewUser(0, "", "", "", "", validateReq.Token, "")
}

func (m Mapper) ToDomainRole(dtoRole string) user.Role {
	return user.Role(dtoRole)
}

func (m Mapper) ToDTORole(domainRole user.Role) string {
	return string(domainRole)
}
