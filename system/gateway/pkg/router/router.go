package router

import (
	access "access/pkg"
	"gateway/pkg/middleware"
	"github.com/gin-gonic/gin"
	order "order/pkg"
)

func Access() *gin.Engine {
	engine := gin.Default()
	group := engine.Group("v1")

	group.POST("/login", access.Login)
	return engine
}

func Order() *gin.Engine {
	engine := gin.Default()
	group := engine.Group("v1")
	group.Use(middleware.JwtAuth)
	group.Use(middleware.AuthenticatedApiRateLimiter)

	group.POST("/orders", order.CreateOrder)
	group.POST("/orders/report", order.CreateOrderReport)
	return engine
}
