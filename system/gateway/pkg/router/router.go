package router

import (
	access "access/pkg"
	"gateway/internal/router/middleware"
	"github.com/gin-gonic/gin"
	order "order/pkg"
)

func Access() *gin.Engine {
	engine := gin.Default()
	group := engine.Group("v1")
	group.Use(middleware.ClientIpRateLimiter)

	group.POST("/login", access.Login)
	return engine
}

func Order() *gin.Engine {
	engine := gin.Default()
	v1 := engine.Group("v1")
	v1.Use(middleware.JwtAuth)

	orders := v1.Group("orders")
	orders.Use(middleware.UserIdRateLimiter)
	orders.POST("", order.CreateOrder)

	reports := orders.Group("reports")
	reports.Use(middleware.TimeRateLimiter)
	reports.POST("", order.CreateOrderReport)
	return engine
}
