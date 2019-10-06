package domain

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
)

// PageServiceServerImpl implements PageServiceServer interface
type PageServiceServerImpl struct {
}

//NewPageServiceServerImpl returns the pointer to the implementation.
func NewPageServiceServerImpl() *PageServiceServerImpl {
	return &PageServiceServerImpl{}
}

// Create function
func (serviceImpl *PageServiceServerImpl) Create(ctx context.Context, p *Page) (*CreatePageResponse, error) {
	id, err := CreatePage(p)
	if err != nil {
		return &CreatePageResponse{
			CreatedPageId: 0,
			Error:         &Error{Message: err.Error()},
		}, err
	}
	return &CreatePageResponse{
		CreatedPageId: int64(id),
		Error:         nil,
	}, nil
}

// GetAll function
func (serviceImpl *PageServiceServerImpl) GetAll(req *empty.Empty, stream PageService_GetAllServer) error {
	pages, err := GetPages()
	if err != nil {
		return err
	}
	for _, p := range pages {
		if err := stream.Send(p); err != nil {
			return err
		}
	}
	return nil
}

// GetOne function
func (serviceImpl *PageServiceServerImpl) GetOne(ctx context.Context, req *GetOnePageRequest) (*Page, error) {
	page, err := GetPage(req.Id)
	if err != nil {
		return nil, err
	}
	return page, nil
}
