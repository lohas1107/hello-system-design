package api

import (
	"e2e/mock"
	"gateway/pkg/middleware"
	"testing"
)

func GivenUserLoggedIn(t *testing.T) {
	accessToken := (*Login(t).Data()).AccessToken
	mock.SetAccessToken(t, accessToken)
}

func GivenUserIp(t *testing.T, ip string) {
	mock.SetClientIp(t, ip)
}

func GivenLastRequestedAt(path string, requestedAt int64) {
	middleware.RateLimitMap[path] = requestedAt
}
