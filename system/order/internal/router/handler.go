package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Engine() *gin.Engine {
	engine := gin.Default()
	v1 := engine.Group("v1", JwtAuth, AuthenticatedApiRateLimiter)
	v1.POST("/orders", CreateOrder)
	v1.POST("/orders/report", CreateOrderReport)
	return engine
}

var RateLimitMap = map[string]int8{}

func JwtAuth(context *gin.Context) {
	token := context.GetHeader("Authentication")
	context.Set("userId", token)
}

func AuthenticatedApiRateLimiter(context *gin.Context) {
	userID := context.GetString("userId")
	RateLimitMap[userID] += 1
	if RateLimitMap[userID] > 5 {
		context.AbortWithStatus(http.StatusTooManyRequests)
	}
}

func CreateOrder(context *gin.Context) {
	context.Status(http.StatusCreated)
}

func CreateOrderReport(context *gin.Context) {
	context.Status(http.StatusAccepted)
}
