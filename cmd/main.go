package main

import (
	"context"
	"log"
	"net/http"
	"time"

	api "github.com/Avon11/Chotu-Go/internal/handler"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client

func init() {
	log.Println("Initializing application...")

	// Initialize Redis client
	log.Println("Connecting to Redis...")
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Test Redis connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Println("Failed to connect to Redis: %v", err)
	}

}

func main() {
	router, err := api.SetupAPIHandler(rdb)
	if err != nil {
		log.Fatalf("Failed to setup API handler: %v", err)
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Close Redis connection
	defer rdb.Close()

	srv.ListenAndServe()
	log.Println("Server exiting")
}
