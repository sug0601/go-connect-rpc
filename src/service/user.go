package service

import (
	"context"
	"fmt"

	examplev1 "example.com/gen/go/proto"
	"example.com/src/repository"
	"github.com/bufbuild/connect-go"
)

type UserServer struct {
	userRepo repository.UserRepository
}

func NewUserServer(repo repository.UserRepository) *UserServer {
	return &UserServer{userRepo: repo}
}

func (s *UserServer) CreateUser(
	ctx context.Context,
	req *connect.Request[examplev1.CreateUserRequest],
) (*connect.Response[examplev1.CreateUserResponse], error) {

	existingUser, err := s.userRepo.FindByEmail(ctx, req.Msg.Email)
	if err == nil && existingUser != nil {
		return nil, connect.NewError(connect.CodeAlreadyExists, fmt.Errorf("email already exists"))
	}

	user, err := s.userRepo.Insert(ctx, req.Msg.Name, req.Msg.Email)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := connect.NewResponse(&examplev1.CreateUserResponse{
		UserId:  user.ID,
		Message: fmt.Sprintf("User %s created successfully!", user.Name),
	})
	return res, nil
}

func (s *UserServer) GetUser(
	ctx context.Context,
	req *connect.Request[examplev1.GetUserRequest],
) (*connect.Response[examplev1.GetUserResponse], error) {

	user, err := s.userRepo.FindByID(ctx, req.Msg.UserId)
	if err != nil {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("user not found"))
	}

	res := connect.NewResponse(&examplev1.GetUserResponse{
		UserId: user.ID,
		Name:   user.Name,
		Email:  user.Email,
	})
	return res, nil
}
