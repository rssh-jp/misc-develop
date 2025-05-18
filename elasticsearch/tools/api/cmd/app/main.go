package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"elasticsearch-search-app/internal/infrastructure/router"
	// 環境変数をロードするため (例: godotenv)
	// _ "github.com/joho/godotenv/autoload"
)

func main() {
	// 環境変数 ELASTICSEARCH_URLS が設定されているか確認
	if os.Getenv("ELASTICSEARCH_URLS") == "" {
		log.Println("Warning: ELASTICSEARCH_URLS environment variable is not set. Defaulting to http://localhost:9200")
		os.Setenv("ELASTICSEARCH_URLS", "http://localhost:9200") // デフォルト値
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // デフォルトポート
	}


	log.Printf("Starting server on :%s...", port)

	r := router.NewRouter() // ルーターの初期化 (依存関係の注入)

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on :%s: %v\n", port, err)
	}
}
