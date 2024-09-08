package router

import "github.com/gin-gonic/gin"

func Engine() *gin.Engine {
	engine := gin.Default()
	v1 := engine.Group("v1")
	v1.POST("/login")
	return engine
}
