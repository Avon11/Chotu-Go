package api

import (
	"github.com/Avon11/Chotu-Go/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Handler struct {
	Service *service.ShortCodeService
}

func NewHandler(service *service.ShortCodeService) *Handler {
	return &Handler{
		Service: service,
	}
}

func SetupAPIHandler(rdb *redis.Client) (*gin.Engine, error) {
	// Initialize the Gin router
	r := gin.Default()
	r.Use(cors.Default())

	service := service.NewCodeService(rdb)

	handler := NewHandler(service)

	r.GET("/get-url", handler.GetUrl)
	r.POST("/post-url", handler.PostUrl)

	return r, nil
}
