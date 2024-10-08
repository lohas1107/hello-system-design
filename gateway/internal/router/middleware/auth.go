package middleware

import (
	"gateway/internal/router/context"
	"github.com/gin-gonic/gin"
)

func JwtAuth(ctx *gin.Context) {
	token := context.Authentication(ctx)
	context.SetUserID(ctx, token)
}
