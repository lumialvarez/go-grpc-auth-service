package auth

import (
	"context"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/pb"
	repositoryUser "github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/repository/postgresql/user"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/utils"
	"net/http"
)

type Handler struct {
	Repository repositoryUser.Repository
	Jwt        utils.JwtWrapper
	pb.UnimplementedAuthServiceServer
}

func (s *Handler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
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

func (s *Handler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
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

func (s *Handler) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	claims, err := s.Jwt.ValidateToken(req.Token)

	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	user, error := s.Repository.GetByEmail(claims.Email)
	if error != nil {
		return &pb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: user.Id,
	}, nil
}
