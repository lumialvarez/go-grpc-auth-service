package devapi

import (
	notificationPublisher "github.com/lumialvarez/go-common-tools/service/rabbitmq/notification"
	"github.com/lumialvarez/go-grpc-auth-service/src/cmd/devapi/config"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/consumers/notification"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/pb"
	errorGrpc "github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/error"
	repositoryUser "github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/repository/postgresql/user"
	serviceJwtUser "github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/service/jwt/user"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user/usecase/create_notification"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user/usecase/current"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user/usecase/list"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user/usecase/login"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user/usecase/read_notification"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user/usecase/register"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user/usecase/update"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user/usecase/validate"
)

type DependenciesContainer struct {
	AuthService pb.AuthServiceServer
}

func LoadDependencies(config config.Config) DependenciesContainer {
	userRepository := repositoryUser.Init(config)
	publishNotificationService := notificationPublisher.Init()

	jwtService := serviceJwtUser.Service{
		SecretKey:       config.JwtSecretKey,
		Issuer:          config.JwtIssuer,
		ExpirationHours: config.JwtExpirationHours,
	}

	userCaseRegister := register.NewUseCaseRegisterUser(&userRepository)
	useCaseLogin := login.NewUseCaseLoginUser(&userRepository, &jwtService, &publishNotificationService)
	useCaseValidate := validate.NewUseCaseValidateUser(&userRepository, &jwtService)
	useCaseList := list.NewUseCaseListUser(&userRepository)
	useCaseUpdate := update.NewUseCaseUpdateUser(&userRepository)
	useCaseCurrent := current.NewUseCaseCurrentUser(&userRepository, &jwtService)
	useCaseCreateNotification := create_notification.NewUseCaseCreateNotification(&userRepository)
	useCaseReadNotification := read_notification.NewUseCaseReadNotification(&userRepository)
	apiResponseProvider := errorGrpc.NewAPIResponseProvider()

	s := auth.NewHandler(userCaseRegister, useCaseLogin, useCaseValidate, useCaseList, useCaseUpdate, useCaseCurrent, useCaseReadNotification, apiResponseProvider)

	notificationConsumer := notification.NewConsumer(useCaseCreateNotification)
	go func() {
		notificationConsumer.Init(config)
	}()

	return DependenciesContainer{
		AuthService: &s,
	}
}
