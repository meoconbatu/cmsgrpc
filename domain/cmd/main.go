package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	_ "github.com/joho/godotenv/autoload"
	"github.com/meoconbatu/cmsgrpc/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	grpcAddress := os.Getenv("GRPC_ADDRESS")
	restAddress := os.Getenv("REST_ADDRESS")

	netListener := getNetListener(grpcAddress)
	go func() {
		err := startGRPCServer(netListener)
		if err != nil {
			log.Fatalf("failed to start gRPC server: %v", err)
		}
	}()

	go func() {
		err := startRESTServer(restAddress, fmt.Sprintf("%s:%s", grpcAddress))
		if err != nil {
			log.Fatalf("failed to start gRPC server: %v", err)
		}
	}()

	select {}
}
func getNetListener(address string) net.Listener {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(fmt.Sprintf("failed to listen: %v", err))
	}
	return lis
}
func startGRPCServer(netListener net.Listener) error {
	creds, err := credentials.NewServerTLSFromFile("cert/server.crt", "cert/server.key")
	if err != nil {
		log.Fatalf("could not load TLS keys: %s", err)
	}
	// Create an array of gRPC options with the credentials
	opts := []grpc.ServerOption{grpc.Creds(creds)}

	gRPCServer := grpc.NewServer(opts...)
	repositoryServiceImpl := domain.NewPageServiceServerImpl()

	domain.RegisterPageServiceServer(gRPCServer, repositoryServiceImpl)

	log.Printf("starting HTTP/2 gRPC server on %s", netListener.Addr())
	if err := gRPCServer.Serve(netListener); err != nil {
		return fmt.Errorf("failed to serve: %s", err)
	}
	return nil
}
func startRESTServer(address, grpcAddress string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	creds, err := credentials.NewClientTLSFromFile("cert/server.crt", "")
	if err != nil {
		return fmt.Errorf("could not load TLS certificate: %s", err)
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}

	mux := runtime.NewServeMux(runtime.WithDisablePathLengthFallback())

	err = domain.RegisterPageServiceHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return fmt.Errorf("could not register service domain: %s", err)
	}
	log.Printf("starting HTTP/1.1 REST server on %s", address)
	http.ListenAndServe(address, mux)

	return nil
}
