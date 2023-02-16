package auth

import (
	"context"
	"github.com/lumialvarez/go-common-tools/validations"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/mapper"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/pb"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
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

type UseCaseList interface {
	Execute(ctx context.Context, id int64, userName string) (*[]user.User, error)
}

type UseCaseUpdate interface {
	Execute(ctx context.Context, domainUser *user.User) (*user.User, error)
}

type UseCaseCurrent interface {
	Execute(ctx context.Context, token string) (*user.User, error)
}

type UseCaseReadNotification interface {
	Execute(ctx context.Context, userId int64, notificationId int64) error
}

type ApiResponseProvider interface {
	ToAPIResponse(err error) error
}

type Handler struct {
	useCaseRegister         UseCaseRegister
	useCaseLogin            UseCaseLogin
	useCaseValidate         UseCaseValidate
	useCaseList             UseCaseList
	useCaseUpdate           UseCaseUpdate
	useCaseCurrent          UseCaseCurrent
	useCaseReadNotification UseCaseReadNotification
	apiResponseProvider     ApiResponseProvider
	mapper.Mapper
	pb.UnimplementedAuthServiceServer
}

func NewHandler(useCaseRegister UseCaseRegister, useCaseLogin UseCaseLogin, useCaseValidate UseCaseValidate, useCaseList UseCaseList, useCaseUpdate UseCaseUpdate, useCaseCurrent UseCaseCurrent, useCaseReadNotification UseCaseReadNotification, apiResponseProvider ApiResponseProvider) Handler {
	return Handler{useCaseRegister: useCaseRegister, useCaseLogin: useCaseLogin, useCaseValidate: useCaseValidate, useCaseList: useCaseList, useCaseUpdate: useCaseUpdate, useCaseCurrent: useCaseCurrent, useCaseReadNotification: useCaseReadNotification, apiResponseProvider: apiResponseProvider}
}

func (s *Handler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	domainUser := s.ToDomainRegisterRequest(req)

	if len(domainUser.UserName()) <= 2 || domainUser.Role() == "" || !validations.ValidEmail(req.Email) {
		return nil, status.Error(codes.InvalidArgument, "Bad request")
	}

	userCreated, err := s.useCaseRegister.Execute(ctx, domainUser)
	if err != nil {
		return nil, s.apiResponseProvider.ToAPIResponse(err)
	}

	return s.ToDTORegisterResponse(userCreated), nil
}

func (s *Handler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	domainUser, err := s.useCaseLogin.Execute(ctx, s.ToDomainLoginRequest(req))
	if err != nil {
		return nil, s.apiResponseProvider.ToAPIResponse(err)
	}

	return s.ToDTOLoginResponse(domainUser), nil
}

func (s *Handler) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	domainUser, err := s.useCaseValidate.Execute(ctx, s.ToDomainValidateRequest(req))
	if err != nil {
		return nil, s.apiResponseProvider.ToAPIResponse(err)
	}

	return s.ToDTOValidateResponse(domainUser), nil
}

func (s *Handler) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	domainUsers, err := s.useCaseList.Execute(ctx, req.GetUserId(), req.GetUserName())
	if err != nil {
		return nil, s.apiResponseProvider.ToAPIResponse(err)
	}

	return s.ToDTOListResponse(domainUsers), nil
}

func (s *Handler) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	domainUser := s.ToDomainUpdateRequest(req)
	domainUser, err := s.useCaseUpdate.Execute(ctx, domainUser)
	if err != nil {
		return nil, s.apiResponseProvider.ToAPIResponse(err)
	}

	return &pb.UpdateResponse{}, nil
}

func (s *Handler) Current(ctx context.Context, req *pb.CurrentRequest) (*pb.CurrentResponse, error) {
	var token string
	md, _ := metadata.FromIncomingContext(ctx)
	authKeys := md.Get("authorization")

	if len(authKeys) == 0 {
		return nil, status.Error(codes.Unauthenticated, "authorization Metadata not found")
	}
	authorization := authKeys[0]

	piecesToken := strings.Split(authorization, "Bearer ")

	if len(piecesToken) < 2 {
		return nil, status.Error(codes.Unauthenticated, "Invalid authorization Metadata")
	}

	token = piecesToken[1]

	println(token)

	domainUser, err := s.useCaseCurrent.Execute(ctx, token)
	if err != nil {
		return nil, s.apiResponseProvider.ToAPIResponse(err)
	}

	println(domainUser)

	return s.ToDTOCurrentResponse(domainUser), nil
}

func (s *Handler) ReadNotification(ctx context.Context, req *pb.ReadNotificationRequest) (*pb.ReadNotificationResponse, error) {
	userId := req.UserId
	notificationId := req.NotificationId

	err := s.useCaseReadNotification.Execute(ctx, userId, notificationId)
	if err != nil {
		return nil, s.apiResponseProvider.ToAPIResponse(err)
	}
	return &pb.ReadNotificationResponse{}, nil
}