package main

import (
	"fmt"
	"log"
	"net"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/meoconbatu/cmsgrpc/user"
	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("PORT")
	netListener := getNetListener(port)
	gRPCServer := grpc.NewServer()
	repositoryServiceImpl := user.NewServiceServerImpl()
	user.RegisterUserServiceServer(gRPCServer, repositoryServiceImpl)

	if err := gRPCServer.Serve(netListener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
func getNetListener(port string) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(fmt.Sprintf("failed to listen: %v", err))
	}
	return lis
}
