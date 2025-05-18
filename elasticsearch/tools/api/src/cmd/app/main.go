package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"elasticsearch-search-app/internal/infrastructure/router"
)

func main() {
	if os.Getenv("ELASTICSEARCH_URLS") == "" {
		log.Println("Warning: ELASTICSEARCH_URLS environment variable is not set. Defaulting to http://localhost:9200")
		os.Setenv("ELASTICSEARCH_URLS", "http://localhost:9200")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on :%s...", port)

	r := router.NewRouter()

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// 終了シグナルを監視するゴルーチン
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on :%s: %v\n", port, err)
		}
	}()

	// シグナルを受信すると graceful shutdown を開始
	<-stop
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting.")
}
