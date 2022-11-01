package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/havus/go-webhook-server/service"
)

type RequestDataHandler interface {
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	Post(c *gin.Context)
	Put(c *gin.Context)
	Delete(c *gin.Context)
}

type RequestDataHandlerImpl struct {
	requestDataService service.RequestDataService
}

func NewRequestDataHandler(requestDataService service.RequestDataService) *RequestDataHandlerImpl {
	return &RequestDataHandlerImpl{
		requestDataService: requestDataService,
	}
}

func (handler *RequestDataHandlerImpl) Post(c *gin.Context) {
	err := handler.requestDataService.Create(c, "POST")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (handler *RequestDataHandlerImpl) GetAll(c *gin.Context) {
	urlQuery 			:= c.Request.URL.Query()
	urlQueryMinId := urlQuery.Get("min_id")
	urlQueryMaxId := urlQuery.Get("max_id")

	response, err := handler.requestDataService.GetAllByAccountId(c, urlQueryMinId, urlQueryMaxId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}
