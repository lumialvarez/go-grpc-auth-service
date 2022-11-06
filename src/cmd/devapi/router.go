package devapi

import (
	"github.com/lumialvarez/go-grpc-auth-service/src/cmd/devapi/config"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/pb"
	"google.golang.org/grpc"
)

func ConfigureServers(grpcServer *grpc.Server, config config.Config) {

	handlers := LoadDependencies(config)

	registerServers(grpcServer, handlers)
}

func registerServers(grpcServer *grpc.Server, handlers DependenciesContainer) {
	pb.RegisterAuthServiceServer(grpcServer, handlers.AuthService)
}
