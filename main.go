package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/havus/go-webhook-server/handler"
	"github.com/havus/go-webhook-server/repository"
	"github.com/havus/go-webhook-server/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    panic(err)
  }

  router := gin.Default()
  router.Use(cors.Default())

  router.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  });

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
	v1 	:= api.Group("/v1")

  request_data_repository := repository.NewRequestDataRepository(db_client.Database(os.Getenv("DATABASE_NAME")))
  request_data_service := service.NewRequestDataService(request_data_repository)
  request_data_handler := handler.NewRequestDataHandler(request_data_service)

  v1.POST("/:account_id/receive", request_data_handler.Post)
  v1.GET("/admin/:account_id/requests", request_data_handler.GetAll)

  router.Run(":" + os.Getenv("PORT"))
}
