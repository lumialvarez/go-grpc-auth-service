package devapi

import (
	"fmt"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/pb"
	repositoryUser "github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/repository/postgresql/user"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/utils"
	"google.golang.org/grpc"
	"net"

	//"fmt"
	"github.com/lumialvarez/go-grpc-auth-service/src/cmd/devapi/config"
	"log"
	//"net"
)

func Start() {
	config, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	userRepository := repositoryUser.Init(config)

	jwt := utils.JwtWrapper{
		SecretKey:       config.JWTSecretKey,
		Issuer:          "go-grpc-auth-service",
		ExpirationHours: 24 * 365,
	}

	lis, err := net.Listen("tcp", config.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth Svc on", config.Port)

	s := auth.Server_borrar{
		Repository: userRepository,
		Jwt:        jwt,
	}

	/*validate.NewUseCaseValidateUser(userRepository, &serviceJwtUser.Service{})
	s := auth.NewHandler()*/

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
