package handler

import "github.com/gin-gonic/gin"

type RequestDataHandler interface {
	Get(c *gin.Context)
	Post(c *gin.Context)
	Put(c *gin.Context)
	Delete(c *gin.Context)
}
