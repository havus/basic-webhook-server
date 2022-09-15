package handler

import "github.com/gin-gonic/gin"

type RequestDataHandler interface {
	Create(c *gin.Context)
}
