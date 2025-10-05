package service_test

import (
	"context"
	"testing"

	"example.com/gen/go/proto"
	"example.com/testutil"
	"github.com/bufbuild/connect-go"
)

func TestHelloService_E2E(t *testing.T) {
	client := testutil.NewHelloClient(t)
	ctx := context.Background()

	t.Run("初回挨拶", func(t *testing.T) {
		req := connect.NewRequest(&proto.HelloRequest{Name: "Alice"})
		res, err := client.SayHello(ctx, req)
		
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if res.Msg.Message != "Hello, Alice!" {
			t.Errorf("got %q, want %q", res.Msg.Message, "Hello, Alice!")
		}
	})

	t.Run("2回目挨拶", func(t *testing.T) {
		req := connect.NewRequest(&proto.HelloRequest{Name: "Alice"})
		res, err := client.SayHello(ctx, req)
		
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if res.Msg.Message != "Hello, Alice!" {
			t.Errorf("got %q, want %q", res.Msg.Message, "Hello, Alice!")
		}
	})

	t.Run("空文字エラー", func(t *testing.T) {
		req := connect.NewRequest(&proto.HelloRequest{Name: ""})
		_, err := client.SayHello(ctx, req)
		
		if err == nil {
			t.Error("expected validation error but got none")
		}
		if connect.CodeOf(err) != connect.CodeInvalidArgument {
			t.Errorf("expected CodeInvalidArgument, got %v", connect.CodeOf(err))
		}
	})

	t.Run("特殊文字エラー", func(t *testing.T) {
		req := connect.NewRequest(&proto.HelloRequest{Name: "Alice@#$"})
		_, err := client.SayHello(ctx, req)
		
		if err == nil {
			t.Error("expected validation error but got none")
		}
		if connect.CodeOf(err) != connect.CodeInvalidArgument {
			t.Errorf("expected CodeInvalidArgument, got %v", connect.CodeOf(err))
		}
	})

	t.Run("スペース含む名前", func(t *testing.T) {
		req := connect.NewRequest(&proto.HelloRequest{Name: "Alice Johnson"})
		res, err := client.SayHello(ctx, req)
		
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if res.Msg.Message != "Hello, Alice Johnson!" {
			t.Errorf("got %q, want %q", res.Msg.Message, "Hello, Alice Johnson!")
		}
	})
}