package main

import (
	"log"
	"net/http"

	"example.com/src/app"
	"example.com/src/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	if cfg.Debug {
		log.Printf("Configuration loaded: %+v", cfg)
	}

	// アプリケーション初期化
	a := app.Initialize(cfg.DatabaseDSN)
	defer a.Close()

	// ハンドラー登録
	mux := http.NewServeMux()
	a.RegisterHandlers(mux)

	// サーバー起動
	log.Printf("Server listening on %s (ENV: %s)", cfg.ServerAddress(), cfg.Environment)
	log.Fatal(http.ListenAndServe(cfg.ServerAddress(), mux))
}
