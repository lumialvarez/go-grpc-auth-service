package devapi

import (
	"github.com/lumialvarez/go-grpc-auth-service/src/cmd/devapi/config"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/pb"
	repositoryUser "github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/repository/postgresql/user"
	serviceJwtUser "github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/service/jwt/user"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/utils"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user/usecase/login"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user/usecase/register"
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

	jwt := utils.JwtWrapper{
		SecretKey:       config.JWTSecretKey,
		Issuer:          "go-grpc-auth-service",
		ExpirationHours: 24 * 365,
	}

	registerUserCase := register.NewUseCaseRegisterUser(&userRepository, &jwtService)
	useCaseLogin := login.NewUseCaseLoginUser(&userRepository, &jwtService)

	/*s := auth.Handler{
		Repository: userRepository,
		Jwt:        jwt,
	}*/
	s := auth.NewHandler(registerUserCase, useCaseLogin, userRepository, jwt)

	return DependenciesContainer{
		AuthService: &s,
	}
}
