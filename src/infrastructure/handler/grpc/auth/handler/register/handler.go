package register

import (
	"context"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/pb"
	repositoryUser "github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/repository/postgresql/user"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/utils"
	"net/http"
)

type HandlerRegister struct {
	Repository repositoryUser.Repository
	Jwt        utils.JwtWrapper
}

func (s *HandlerRegister) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user, error := s.Repository.GetByEmail(req.GetEmail())
	if error == nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "E-Mail already exists",
		}, nil
	}

	user.Email = req.Email
	user.Password = utils.HashPassword(req.Password)

	s.Repository.Save(user)

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}
