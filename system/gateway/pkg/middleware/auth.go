package middleware

import "github.com/gin-gonic/gin"

func JwtAuth(context *gin.Context) {
	token := context.GetHeader("Authentication")
	context.Set("userId", token)
}
