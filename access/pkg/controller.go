package pkg

import (
	"gateway/pkg/io"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(context *gin.Context) {
	context.JSON(
		http.StatusOK,
		io.LoginResponse{
			AccessToken: "123",
		},
	)
}
