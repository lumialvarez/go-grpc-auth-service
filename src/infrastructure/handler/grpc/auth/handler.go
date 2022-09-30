package auth

import (
	"context"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/pb"
)

type UseCaseValidate interface {
	//Execute(ctx gin.Context, routes *[]domainRoute.Route) ([]domainRoute.Route, error)
}

type Handler struct {
	useCaseValidate UseCaseValidate
}

func NewHandler(useCaseValidate UseCaseValidate) *Handler {
	return &Handler{useCaseValidate: useCaseValidate}
}

func (s *Handler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return nil, nil
}

func (s *Handler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return nil, nil
}

func (s *Handler) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	return nil, nil
}
