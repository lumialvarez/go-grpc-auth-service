package devapi

import (
	"github.com/lumialvarez/go-grpc-auth-service/src/cmd/devapi/config"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/pb"
	errorGrpc "github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/error"
	repositoryUser "github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/repository/postgresql/user"
	serviceJwtUser "github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/service/jwt/user"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user/usecase/list"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user/usecase/login"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user/usecase/register"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user/usecase/update"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user/usecase/validate"
)

type DependenciesContainer struct {
	AuthService pb.AuthServiceServer
}

func LoadDependencies(config config.Config) DependenciesContainer {
	userRepository := repositoryUser.Init(config)

	jwtService := serviceJwtUser.Service{
		SecretKey:       config.JWTSecretKey,
		Issuer:          "go-grpc-auth-service",
		ExpirationHours: 24 * 365,
	}

	userCaseRegister := register.NewUseCaseRegisterUser(&userRepository)
	useCaseLogin := login.NewUseCaseLoginUser(&userRepository, &jwtService)
	useCaseValidate := validate.NewUseCaseValidateUser(&userRepository, &jwtService)
	useCaseList := list.NewUseCaseListUser(&userRepository)
	useCaseUpdate := update.NewUseCaseUpdateUser(&userRepository)
	apiResponseProvider := errorGrpc.NewAPIResponseProvider()

	s := auth.NewHandler(userCaseRegister, useCaseLogin, useCaseValidate, useCaseList, useCaseUpdate, apiResponseProvider)

	return DependenciesContainer{
		AuthService: &s,
	}
}
