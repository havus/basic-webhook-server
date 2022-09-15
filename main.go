package main

import (
	"basic-webhook-server/handler"
	"basic-webhook-server/repository"
	"basic-webhook-server/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()

  router.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  });

  // VERSIONING
	api := router.Group("/api")
	v1 	:= api.Group("/v1")

  request_data_repository := repository.NewRequestDataRepository()
  request_data_service := service.NewRequestDataService(request_data_repository)
  request_data_handler := handler.NewRequestDataHandler(request_data_service)

  v1.POST("/request/:account_id", request_data_handler.Create)

  router.Run(":3000")
}
