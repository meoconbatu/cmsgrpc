package main

import (
	"fmt"
	"log"
	"net"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/meoconbatu/cmsgrpc/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	netListener := getNetListener(host, port)

	creds, err := credentials.NewServerTLSFromFile("cert/server.crt", "cert/server.key")
	if err != nil {
		log.Fatalf("could not load TLS keys: %s", err)
	}
	// Create an array of gRPC options with the credentials
	opts := []grpc.ServerOption{grpc.Creds(creds)}

	gRPCServer := grpc.NewServer(opts...)
	repositoryServiceImpl := domain.NewPageServiceServerImpl()

	domain.RegisterPageServiceServer(gRPCServer, repositoryServiceImpl)

	if err := gRPCServer.Serve(netListener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
func getNetListener(host, port string) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(fmt.Sprintf("failed to listen: %v", err))
	}
	return lis
}
