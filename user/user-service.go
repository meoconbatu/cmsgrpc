package user

import "context"

// ServiceServerImpl struct
type ServiceServerImpl struct {
}

// NewServiceServerImpl return the pointer to the ServiceServerImpl struct
func NewServiceServerImpl() *ServiceServerImpl {
	return &ServiceServerImpl{}
}

// AuthenticateUser function
func (serviceImpl *ServiceServerImpl) AuthenticateUser(ctx context.Context, req *AuthenticateUserRequest) (*AuthenticateUserResponse, error) {
	err := AuthenticateUser(req.UserName, req.Password)
	return &AuthenticateUserResponse{}, err
}
