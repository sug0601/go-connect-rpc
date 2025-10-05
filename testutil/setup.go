package testutil

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"example.com/gen/go/proto/protoconnect"
	"example.com/src/app"
	"github.com/uptrace/bun"
)

func SetupTestServer(t *testing.T) (*app.App, *httptest.Server) {
	t.Helper()

	dsn := os.Getenv("TEST_DB_DSN")
	if dsn == "" {
		dsn = "postgres://user:pass@localhost:5433/connect?sslmode=disable"
	}

	// アプリケーション初期化
	a := app.Initialize(dsn)
	t.Cleanup(a.Close)

	// テーブルクリーンアップ
	CleanupDB(t, a.DB)

	// HTTPサーバー
	mux := http.NewServeMux()
	a.RegisterHandlers(mux)

	ts := httptest.NewServer(mux)
	t.Cleanup(ts.Close)

	return a, ts
}

func NewHelloClient(t *testing.T) protoconnect.HelloServiceClient {
	t.Helper()
	_, ts := SetupTestServer(t)
	return protoconnect.NewHelloServiceClient(http.DefaultClient, ts.URL)
}

func NewUserClient(t *testing.T) protoconnect.UserServiceClient {
	t.Helper()
	_, ts := SetupTestServer(t)
	return protoconnect.NewUserServiceClient(http.DefaultClient, ts.URL)
}

func CleanupDB(t *testing.T, db *bun.DB) {
	t.Helper()
	tables := []string{"greetings", "users"}
	for _, table := range tables {
		_, err := db.Exec("TRUNCATE TABLE " + table + " CASCADE")
		if err != nil {
			t.Logf("Warning: failed to truncate %s: %v", table, err)
		}
	}
}