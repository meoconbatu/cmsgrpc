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
	if err != nil {
		return &AuthenticateUserResponse{Error: &Error{Code: err.Error(), Message: err.Error()}}, nil
	}
	return &AuthenticateUserResponse{}, nil
}

// NewUser function
func (serviceImpl *ServiceServerImpl) NewUser(ctx context.Context, u *User) (*NewUserResponse, error) {
	err := NewUser(u.UserName, u.Password)
	if err != nil {
		return &NewUserResponse{Error: &Error{Code: err.Error(), Message: err.Error()}}, nil
	}
	return &NewUserResponse{}, nil
}
