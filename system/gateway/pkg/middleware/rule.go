package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var RateLimitMap = map[string]int8{}

func AuthenticatedApiRateLimiter(context *gin.Context) {
	userID := context.GetString("userId")
	RateLimitMap[userID] += 1
	if RateLimitMap[userID] > 5 {
		context.AbortWithStatus(http.StatusTooManyRequests)
	}
}
