package view

import (
	"context"
	"errors"

	"google.golang.org/grpc"
)

// AuthenticateUser authenticate user
func AuthenticateUser(username, password string) error {
	conn, err := grpc.Dial(serverUserAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := NewUserServiceClient(conn)
	resp, err := client.AuthenticateUser(context.Background(), &AuthenticateUserRequest{
		UserName: username,
		Password: password,
	})
	if err != nil {
		return err
	}
	if resp.Error != nil {
		return errors.New(resp.Error.Message)
	}
	return nil
}

// NewUser create new user
func NewUser(username, password string) error {
	conn, err := grpc.Dial(serverUserAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := NewUserServiceClient(conn)
	resp, err := client.NewUser(context.Background(), &User{
		UserName: username,
		Password: password,
	})
	if err != nil {
		return err
	}
	if resp.Error != nil {
		return errors.New(resp.Error.Message)
	}
	return nil
}
