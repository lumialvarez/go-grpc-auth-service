package auth

import (
	"context"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/mapper"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/pb"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user"
	"net/http"
)

type UseCaseRegister interface {
	Execute(ctx context.Context, domainUser *user.User) error
}

type UseCaseLogin interface {
	Execute(ctx context.Context, domainUser *user.User) (*user.User, error)
}

type UseCaseValidate interface {
	Execute(ctx context.Context, domainUser *user.User) (*user.User, error)
}

type Handler struct {
	useCaseRegister UseCaseRegister
	useCaseLogin    UseCaseLogin
	useCaseValidate UseCaseValidate
	mapper.Mapper
	pb.UnimplementedAuthServiceServer
}

func NewHandler(useCaseRegister UseCaseRegister, useCaseLogin UseCaseLogin, useCaseValidate UseCaseValidate) Handler {
	return Handler{useCaseRegister: useCaseRegister, useCaseLogin: useCaseLogin, useCaseValidate: useCaseValidate}
}

func (s *Handler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	error := s.useCaseRegister.Execute(ctx, s.ToDomainRegister(req))
	if error != nil {
		return &pb.RegisterResponse{
			Status: http.StatusUnauthorized,
			Error:  error.Error(),
		}, nil
	}

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Handler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	domainUser, error := s.useCaseLogin.Execute(ctx, s.ToDomainLogin(req))
	if error != nil {
		return &pb.LoginResponse{
			Status: http.StatusUnauthorized,
			Error:  error.Error(),
		}, nil
	}

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  domainUser.Token(),
	}, nil
}

func (s *Handler) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	domainUser, error := s.useCaseValidate.Execute(ctx, s.ToDomainValidate(req))
	if error != nil {
		return &pb.ValidateResponse{
			Status: http.StatusUnauthorized,
			Error:  error.Error(),
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: domainUser.Id(),
	}, nil
}
