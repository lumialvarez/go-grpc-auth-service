package login

import (
	"context"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/pb"
	repositoryUser "github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/repository/postgresql/user"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/utils"
	"net/http"
)

type HandlerLogin struct {
	Repository repositoryUser.Repository
	Jwt        utils.JwtWrapper
}

func (s *HandlerLogin) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, error := s.Repository.GetByEmail(req.GetEmail())
	if error != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	match := utils.CheckPasswordHash(req.Password, user.Password)

	if !match {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	token, _ := s.Jwt.GenerateToken(user)

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}
