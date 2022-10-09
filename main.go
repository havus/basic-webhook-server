package main

import (
	"basic-webhook-server/handler"
	"basic-webhook-server/repository"
	"basic-webhook-server/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	db_client := connectMongoDb()

	// disconnect mongodb sigterm
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		disconnectMongoDb(db_client)
		os.Exit(1)
	}()

	// VERSIONING
	api := router.Group("/api")
	v1 := api.Group("/v1")

	request_data_repository := repository.NewRequestDataRepository(db_client.Database(os.Getenv("DATABASE_NAME")))
	request_data_service := service.NewRequestDataService(request_data_repository)
	request_data_handler := handler.NewRequestDataHandler(request_data_service)

	v1.POST("/request/:account_id", request_data_handler.Post)

	router.Run(":3000")
}
