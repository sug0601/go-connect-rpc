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
	name := req.Msg.GetName()

	if name == "" {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			fmt.Errorf("validation error: name must not be empty (min_len=1)"),
		)
	}

	// DB チェック
	exists, err := s.greetingRepo.Exists(ctx, name)
	if err != nil {
		return nil, err
	}

	if !exists {
		if err := s.greetingRepo.Insert(ctx, name); err != nil {
			return nil, err
		}
	}

	// レスポンス
	res := connect.NewResponse(&examplev1.HelloResponse{
		Message: fmt.Sprintf("Hello, %s!", name),
	})
	return res, nil
}
