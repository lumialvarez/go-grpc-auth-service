package mapper

import (
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/pb"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user"
)

type Mapper struct {
}

func (m Mapper) ToDomainRegisterRequest(registerReq *pb.RegisterRequest) *user.User {
	return user.NewUser(0, registerReq.Name, registerReq.UserName, registerReq.Email, registerReq.Password, "", m.toDomainRole(registerReq.Role))
}

func (m Mapper) ToDTORegisterResponse(domainUser *user.User) *pb.RegisterResponse {
	return &pb.RegisterResponse{UserId: domainUser.Id()}
}

func (m Mapper) ToDomainLoginRequest(loginReq *pb.LoginRequest) *user.User {
	return user.NewUser(0, "", loginReq.UserName, "", loginReq.Password, "", "")
}

func (m Mapper) ToDTOLoginResponse(domainUser *user.User) *pb.LoginResponse {
	return &pb.LoginResponse{
		Token:    domainUser.Token(),
		UserId:   domainUser.Id(),
		UserName: domainUser.UserName(),
		Role:     m.toDTORole(domainUser.Role()),
	}
}

func (m Mapper) ToDomainValidateRequest(validateReq *pb.ValidateRequest) *user.User {
	return user.NewUser(0, "", "", "", "", validateReq.Token, "")
}

func (m Mapper) ToDTOValidateResponse(domainUser *user.User) *pb.ValidateResponse {
	return &pb.ValidateResponse{
		UserId:   domainUser.Id(),
		UserName: domainUser.UserName(),
		Role:     m.toDTORole(domainUser.Role()),
	}
}

func (m Mapper) toDomainRole(dtoRole string) user.Role {
	return user.Role(dtoRole)
}

func (m Mapper) toDTORole(domainRole user.Role) string {
	return string(domainRole)
}
