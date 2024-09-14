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

func Login[T *io.LoginResponse](t *testing.T) io.Response[T] {
	response := apitest.New().
		Handler(router.Access()).
		Post("/v1/login").
		Expect(t)

	return io.Response[T]{
		Response: response,
		Result:   response.End(),
	}
}
