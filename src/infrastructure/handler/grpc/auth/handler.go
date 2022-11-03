package auth

import (
	"context"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/mapper"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/pb"
	repositoryUser "github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/repository/postgresql/user"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/utils"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user"
	"net/http"
)

type UseCaseLogin interface {
	Execute(ctx context.Context, domainUser *user.User) (*user.User, error)
}

type Handler struct {
	useCaseLogin UseCaseLogin
	repository   repositoryUser.Repository
	jwt          utils.JwtWrapper
	mapper       mapper.Mapper
	pb.UnimplementedAuthServiceServer
}

func NewHandler(useCaseLogin UseCaseLogin, repository repositoryUser.Repository, jwt utils.JwtWrapper) Handler {
	return Handler{useCaseLogin: useCaseLogin, repository: repository, jwt: jwt, mapper: mapper.Mapper{}}
}

func (s *Handler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	userl, error := s.repository.GetByEmail(req.GetEmail())
	if error == nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "E-Mail already exists",
		}, nil
	}

	//userl.Email = req.Email
	//userl.Password = utils.HashPassword(req.Password)

	s.repository.Save(userl)

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Handler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	domainUser, error := s.useCaseLogin.Execute(ctx, s.mapper.ToDomainLogin(req))
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
	claims, err := s.jwt.ValidateToken(req.Token)

	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	userl, error := s.repository.GetByEmail(claims.Email)
	if error != nil {
		return &pb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: userl.Id(),
	}, nil
}
