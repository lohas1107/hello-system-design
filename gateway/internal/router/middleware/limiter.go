package middleware

import (
	"gateway/internal/router/context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var (
	RateLimitRule = map[context.Operation]gin.HandlerFunc{
		context.Operation{Method: http.MethodPost, Route: "/v1/login"}:          ClientIpRateLimiter,
		context.Operation{Method: http.MethodPost, Route: "/v1/orders"}:         UserIdRateLimiter,
		context.Operation{Method: http.MethodPost, Route: "/v1/orders/reports"}: TimeRateLimiter,
	}
	OperationLookup = map[context.Actor]map[context.Operation][]time.Time{}
)

func RateLimiter(c *gin.Context) {
	handle, exist := RateLimitRule[context.RequestOperation(c)]
	switch exist {
	case true:
		handle(c)
	default:
		UserIdRateLimiter(c)
	}
}

func UserIdRateLimiter(c *gin.Context) {
	pass := rateLimit(
		context.Actor(context.UserID(c)),
		time.Second,
		context.RequestOperation(c),
		5,
	)
	if !pass {
		c.AbortWithStatus(http.StatusTooManyRequests)
	}
}

func ClientIpRateLimiter(c *gin.Context) {
	pass := rateLimit(
		context.Actor(context.ClientIp(c)),
		time.Second,
		context.RequestOperation(c),
		5,
	)
	if !pass {
		c.AbortWithStatus(http.StatusTooManyRequests)
	}
}

func TimeRateLimiter(c *gin.Context) {
	pass := rateLimit(
		context.Actor(c.Request.URL.Path),
		time.Second,
		context.RequestOperation(c),
		1,
	)
	if !pass {
		c.AbortWithStatus(http.StatusTooManyRequests)
	}
}

func rateLimit(
	actor context.Actor,
	period time.Duration,
	operation context.Operation,
	threshold int,
) bool {
	operatedAt := time.Now().UTC()
	window := operatedAt.Add(-period)

	_, exist := OperationLookup[actor]
	if !exist {
		OperationLookup[actor] = make(map[context.Operation][]time.Time)
	}

	history := OperationLookup[actor][operation]
	history = append(history, operatedAt)

	for i, t := range history {
		if !t.After(window) {
			continue
		}

		history = history[i:]
		OperationLookup[actor][operation] = history

		if len(history) > threshold {
			return false
		}
		break
	}
	return true
}
