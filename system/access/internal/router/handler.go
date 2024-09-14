package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Engine() *gin.Engine {
	engine := gin.Default()
	v1 := engine.Group("v1")
	v1.POST("/login", Login)
	return engine
}

func Login(context *gin.Context) {
	context.JSON(
		http.StatusOK,
		LoginResponse{
			AccessToken: "123",
		},
	)
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}
