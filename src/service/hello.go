package service

import (
	"context"
	"fmt"

	"example.com/src/repository"

	examplev1 "example.com/gen/go/proto"
	"github.com/bufbuild/connect-go"
)

type HelloServer struct {
	greetingRepo repository.GreetingRepository
}

func NewHelloServer(repo repository.GreetingRepository) *HelloServer {
	return &HelloServer{greetingRepo: repo}
}

func (s *HelloServer) SayHello(
	ctx context.Context,
	req *connect.Request[examplev1.HelloRequest],
) (*connect.Response[examplev1.HelloResponse], error) {

	exists, err := s.greetingRepo.Exists(ctx, req.Msg.Name)
	if err != nil {
		return nil, err
	}

	if !exists {
		if err := s.greetingRepo.Insert(ctx, req.Msg.Name); err != nil {
			return nil, err
		}
	}

	res := connect.NewResponse(&examplev1.HelloResponse{
		Message: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	return res, nil
}
