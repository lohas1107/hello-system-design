package api

import (
	"e2e/config"
	"gateway/pkg/io"
	"gateway/pkg/router"
	"github.com/steinfletcher/apitest"
	"testing"
)

func GivenUserLoggedIn(t *testing.T) {
	accessToken := (*Login(t).Data()).AccessToken
	config.SetAccessToken(t, accessToken)
}

func GivenUserIp(t *testing.T, ip string) {
	config.SetClientIp(t, ip)
}

func Login[T *io.LoginResponse](t *testing.T) io.Response[T] {
	result := apitest.New().
		Handler(router.Access()).
		Post("/v1/login").
		Headers(io.HttpHeaders(
			io.ClientIp(config.ClientIp()),
		)).
		Expect(t).
		End()

	return io.Response[T]{Result: result}
}
