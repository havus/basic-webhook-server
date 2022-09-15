package handler

import (
	"basic-webhook-server/handler/request"
	"basic-webhook-server/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestDataHandlerImpl struct {
	requestDataService service.RequestDataService
}

func NewRequestDataHandler(requestDataService service.RequestDataService) *RequestDataHandlerImpl {
	return &RequestDataHandlerImpl{
		requestDataService: requestDataService,
	}
}

func (handler *RequestDataHandlerImpl) Create(c *gin.Context) {
	var requestDataRequest request.RequestDataRequest

	if err := c.ShouldBindJSON(&requestDataRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	request_data_response := handler.requestDataService.Create(c, requestDataRequest)

	c.JSON(http.StatusCreated, gin.H{
		"data": request_data_response,
	})
}
