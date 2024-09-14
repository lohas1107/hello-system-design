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

var IdentityHandler = apitest.New().Handler(router.Router())

func Login(t *testing.T) io.Response[*io.LoginResponse] {
	response := IdentityHandler.Post("/v1/login").Expect(t)
	return io.Response[*io.LoginResponse]{
		Response: response,
		Result:   response.End(),
	}
}
