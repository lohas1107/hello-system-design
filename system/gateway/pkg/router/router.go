package router

import (
	"access/pkg"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	engine := gin.Default()
	v1 := engine.Group("v1")
	v1.POST("/login", pkg.Login)
	return engine
}
