package middleware

import (
	"gateway/pkg/context"
	"github.com/gin-gonic/gin"
	"net/http"
)

var RateLimitMap = map[string]int8{}

func AuthenticatedApiRateLimiter(ctx *gin.Context) {
	userID := context.UserID(ctx)
	RateLimitMap[userID] += 1
	if RateLimitMap[userID] > 5 {
		ctx.AbortWithStatus(http.StatusTooManyRequests)
	}
}

func UnauthenticatedApiRateLimiter(ctx *gin.Context) {
	ip := context.ClientIp(ctx)
	RateLimitMap[ip] += 1
	if RateLimitMap[ip] > 5 {
		ctx.AbortWithStatus(http.StatusTooManyRequests)
	}
}
