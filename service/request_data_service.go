package service

import (
	"basic-webhook-server/handler/response"

	"github.com/gin-gonic/gin"
)

type RequestDataService interface {
	Create(ctx *gin.Context, request_method string) (*response.RequestDataResponse, error)
}
