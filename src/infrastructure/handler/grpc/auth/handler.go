package auth

import (
	"context"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/mapper"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/pb"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/tmp_utils/validations"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/tmp_utils/validations/passwordvalidator"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UseCaseRegister interface {
	Execute(ctx context.Context, domainUser *user.User) (*user.User, error)
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
	domainUser := s.ToDomainRegisterRequest(req)

	err := passwordvalidator.Validate(req.Password, 80)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Password ("+err.Error()+")")
	}

	if len(domainUser.UserName()) <= 2 || domainUser.Role() == "" || !validations.ValidEmail(req.Email) {
		return nil, status.Error(codes.InvalidArgument, "Bad request")
	}

	userCreated, err := s.useCaseRegister.Execute(ctx, domainUser)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	return s.ToDTORegisterResponse(userCreated), nil
}

func (s *Handler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	domainUser, err := s.useCaseLogin.Execute(ctx, s.ToDomainLoginRequest(req))
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	return s.ToDTOLoginResponse(domainUser), nil
}

func (s *Handler) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	domainUser, err := s.useCaseValidate.Execute(ctx, s.ToDomainValidateRequest(req))
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	return s.ToDTOValidateResponse(domainUser), nil
}
