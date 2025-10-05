package service_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"example.com/gen/go/proto"
	"example.com/gen/go/proto/protoconnect"
	"example.com/src/infra"
	"example.com/src/repository"
	"example.com/src/service"
	"github.com/bufbuild/connect-go"
)

func setupTestServer(t *testing.T) protoconnect.HelloServiceClient {
	t.Helper()

	// テスト用DB
	dsn := "postgres://user:pass@localhost:5433/connect?sslmode=disable"
	db := infra.NewPostgresDB(dsn)
	t.Cleanup(func() { db.Close() })

	// テーブルクリーンアップ
	if _, err := db.Exec("TRUNCATE TABLE greetings"); err != nil {
		t.Fatalf("failed to truncate table: %v", err)
	}

	// サービス初期化
	repo := repository.NewBunGreetingRepository(db)
	server := service.NewHelloServer(repo)

	// HTTPサーバー
	mux := http.NewServeMux()
	path, handler := protoconnect.NewHelloServiceHandler(server)
	mux.Handle(path, handler)

	ts := httptest.NewServer(mux)
	t.Cleanup(ts.Close)

	return protoconnect.NewHelloServiceClient(http.DefaultClient, ts.URL)
}

func TestHelloService_E2E(t *testing.T) {
	client := setupTestServer(t)
	ctx := context.Background()

	t.Run("初回挨拶", func(t *testing.T) {
		req := connect.NewRequest(&proto.HelloRequest{Name: "Alice"})
		res, err := client.SayHello(ctx, req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if res == nil || res.Msg == nil {
			t.Fatal("response is nil")
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
		if res == nil || res.Msg == nil {
			t.Fatal("response is nil")
		}
		if res.Msg.Message != "Hello, Alice!" {
			t.Errorf("got %q, want %q", res.Msg.Message, "Hello, Alice!")
		}
	})

	t.Run("空文字エラー", func(t *testing.T) {
		req := connect.NewRequest(&proto.HelloRequest{Name: ""})
		_, err := client.SayHello(ctx, req)

		if err == nil {
			t.Fatal("expected validation error but got none")
		}

		// ログを出す
		t.Logf("error: %v", err)

		code := connect.CodeOf(err)
		if code != connect.CodeInvalidArgument {
			t.Errorf("expected CodeInvalidArgument, got %v", code)
		}

		errMsg := err.Error()
		if !strings.Contains(errMsg, "min_len") && !strings.Contains(errMsg, "validation error") {
			t.Errorf("expected validation error message, got: %v", errMsg)
		}
	})
}
