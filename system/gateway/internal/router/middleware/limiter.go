package middleware

import (
	"gateway/internal/router/context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var RateLimitMap = map[string]int64{}

func UserIdRateLimiter(c *gin.Context) {
	userID := context.UserID(c)
	RateLimitMap[userID] += 1
	if RateLimitMap[userID] > 5 {
		c.AbortWithStatus(http.StatusTooManyRequests)
	}
}

func ClientIpRateLimiter(c *gin.Context) {
	ip := context.ClientIp(c)
	RateLimitMap[ip] += 1
	if RateLimitMap[ip] > 5 {
		c.AbortWithStatus(http.StatusTooManyRequests)
	}
}

func TimeRateLimiter(c *gin.Context) {
	path := c.Request.URL.Path
	last, exists := RateLimitMap[path]
	now := time.Now().UTC().UnixMilli()
	RateLimitMap[path] = now
	if exists {
		if now-last < 1e3 {
			c.AbortWithStatus(http.StatusTooManyRequests)
		}
	}
}
