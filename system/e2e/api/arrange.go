package api

import (
	"e2e/mock"
	"testing"
)

func GivenUserLoggedIn(t *testing.T) {
	accessToken := (*Login(t).Data()).AccessToken
	mock.SetAccessToken(t, accessToken)
}

func GivenUserIp(t *testing.T, ip string) {
	mock.SetClientIp(t, ip)
}
