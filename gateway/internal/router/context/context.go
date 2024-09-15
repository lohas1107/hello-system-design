package context

import (
	"github.com/gin-gonic/gin"
)

func Authentication(ctx *gin.Context) string {
	return ctx.GetHeader("Authentication")
}

func ClientIp(ctx *gin.Context) string {
	ip := ctx.GetHeader("HTTP_CLIENT_IP")
	return ip
}

func SetUserID(ctx *gin.Context, token string) {
	ctx.Set("userId", token)
}

func UserID(ctx *gin.Context) string {
	return ctx.GetString("userId")
}

type Actor string

type Operation struct {
	Method string
	Route  string
}

func RequestOperation(c *gin.Context) Operation {
	return Operation{
		Method: c.Request.Method,
		Route:  c.Request.URL.Path,
	}
}
