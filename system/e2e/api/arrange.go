package api

import (
	env "e2e/mock"
	"gateway/pkg/mock"
	"testing"
)

func GivenUserLoggedIn(t *testing.T) {
	accessToken := (*Login(t).Data()).AccessToken
	env.SetAccessToken(t, accessToken)
}

func GivenUserIp(t *testing.T, ip string) {
	env.SetClientIp(t, ip)
}

func GivenLastRequestedAt(path string, requestedAt int64) {
	mock.SetRequestedAt(path, requestedAt)
}
