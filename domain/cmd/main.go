package main

import (
	"fmt"
	"log"
	"net"

	"github.com/meoconbatu/cmsgrpc/domain"

	"google.golang.org/grpc"
)

func main() {
	netListener := getNetListener(8080)
	gRPCServer := grpc.NewServer()
	repositoryServiceImpl := domain.NewPageServiceServerImpl()
	domain.RegisterPageServiceServer(gRPCServer, repositoryServiceImpl)

	if err := gRPCServer.Serve(netListener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
func getNetListener(port uint) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(fmt.Sprintf("failed to listen: %v", err))
	}
	return lis
}
