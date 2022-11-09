package auth

import (
	"context"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/mapper"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/pb"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/tmp_utils/validations"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/tmp_utils/validations/passwordvalidator"
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
	domainUser := s.ToDomainRegister(req)

	err := passwordvalidator.Validate(req.Password, 80)
	if err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusBadRequest,
			Error:  "Invalid Password (" + err.Error() + ")",
		}, nil
	}

	if len(domainUser.UserName()) <= 2 || domainUser.Role() == "" || !validations.ValidEmail(req.Email) {
		return &pb.RegisterResponse{
			Status: http.StatusBadRequest,
			Error:  "Bad request",
		}, nil
	}

	err = s.useCaseRegister.Execute(ctx, domainUser)
	if err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusUnauthorized,
			Error:  err.Error(),
		}, nil
	}

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil

	//return nil, status.Error(codes.PermissionDenied, "PERMISSION_DENIED_TEXT")
}

func (s *Handler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	domainUser, err := s.useCaseLogin.Execute(ctx, s.ToDomainLogin(req))
	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusUnauthorized,
			Error:  err.Error(),
		}, nil
	}

	role := s.Mapper.ToDTORole(domainUser.Role())

	return &pb.LoginResponse{
		Status:   http.StatusOK,
		Token:    domainUser.Token(),
		UserId:   domainUser.Id(),
		UserName: domainUser.UserName(),
		Role:     role,
	}, nil
}

func (s *Handler) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	domainUser, err := s.useCaseValidate.Execute(ctx, s.ToDomainValidate(req))
	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusUnauthorized,
			Error:  err.Error(),
		}, nil
	}

	return &pb.ValidateResponse{
		Status:   http.StatusOK,
		UserId:   domainUser.Id(),
		UserName: domainUser.UserName(),
		Role:     s.Mapper.ToDTORole(domainUser.Role()),
	}, nil
}
