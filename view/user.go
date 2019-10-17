package view

import (
	"context"

	"google.golang.org/grpc"
)

const serverUserAddress = "localhost:8090"

// AuthenticateUser authenticate user
func AuthenticateUser(username, password string) error {
	conn, err := grpc.Dial(serverUserAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := NewUserServiceClient(conn)
	_, err = client.AuthenticateUser(context.Background(), &AuthenticateUserRequest{
		UserName: username,
		Password: password,
	})
	if err != nil {
		return err
	}
	return nil
}
