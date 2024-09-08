package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Engine() *gin.Engine {
	engine := gin.Default()
	v1 := engine.Group("v1")
	v1.POST("/orders", CreateOrder)
	v1.POST("/orders/report", CreateOrderReport)
	return engine
}

func CreateOrder(context *gin.Context) {
	context.Status(http.StatusCreated)
}

func CreateOrderReport(context *gin.Context) {
	context.Status(http.StatusAccepted)
}
