package view

import (
	"context"
	"io"

	"google.golang.org/grpc"
)

var client PageServiceClient

const serverAddress = "localhost:8080"

// CreatePage return PageServiceClient
func CreatePage(p *Page) (int, error) {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := NewPageServiceClient(conn)
	responseMessage, e := client.Create(context.Background(), p)
	if e != nil || responseMessage.Error != nil {
		return 0, e
	}
	return int(responseMessage.CreatedPageId), nil
}

// GetPages function
func GetPages() ([]*Page, error) {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := NewPageServiceClient(conn)
	stream, err := client.GetAll(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	pages := make([]*Page, 0)
	for {
		page, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		pages = append(pages, page)
	}
	return pages, nil
}

// GetPage function
func GetPage(id string) (*Page, error) {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := NewPageServiceClient(conn)
	page, err := client.GetOne(context.Background(), &GetOnePageRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return page, nil
}
