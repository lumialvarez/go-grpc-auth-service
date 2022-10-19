package devapi

import (
	"fmt"
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

	lis, err := net.Listen("tcp", config.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth Svc on", config.Port)

	/*validate.NewUseCaseValidateUser(userRepository, &serviceJwtUser.Service{})
	s := auth.NewHandler()*/

	grpcServer := grpc.NewServer()

	ConfigureServers(grpcServer, config)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
