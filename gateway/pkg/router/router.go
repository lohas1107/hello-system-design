package router

import (
	access "access/pkg"
	"gateway/internal/router/middleware"
	"github.com/gin-gonic/gin"
	order "order/pkg"
)

func Access() *gin.Engine {
	engine := gin.Default()
	v1 := engine.Group("v1")
	v1.Use(middleware.RateLimiter)

	v1.POST("/login", access.Login)
	return engine
}

func Order() *gin.Engine {
	engine := gin.Default()
	v1 := engine.Group("v1")
	v1.Use(middleware.JwtAuth)
	v1.Use(middleware.RateLimiter)

	v1.POST("orders", order.CreateOrder)
	v1.POST("orders/reports", order.CreateOrderReport)
	return engine
}
