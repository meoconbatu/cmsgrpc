package main

import (
	"fmt"
	"log"
	"net"

	"github.com/meoconbatu/cmsgrpc/user"

	"google.golang.org/grpc"
)

func main() {
	netListener := getNetListener(8090)
	gRPCServer := grpc.NewServer()
	repositoryServiceImpl := user.NewServiceServerImpl()
	user.RegisterUserServiceServer(gRPCServer, repositoryServiceImpl)

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
