package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	api "github.com/Avon11/ShrinkRay/internal/handler"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

var ctx = context.Background()
var rdb *redis.Client

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// Get db connection from env file
	RedisAddr := os.Getenv("RedisAddr")
	RedisPass := os.Getenv("RedisPass")
	RedisDB := cast.ToInt(os.Getenv("RedisDB"))
	log.Println("Initializing application...")

	// Initialize Redis client
	log.Println("Connecting to Redis...")
	rdb = redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		Password: RedisPass,
		DB:       RedisDB,
	})

	// Test Redis connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
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
