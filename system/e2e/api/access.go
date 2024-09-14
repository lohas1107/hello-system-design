package api

import (
	"access/pkg"
	"e2e/config"
	io "gateway/pkg"
	"github.com/steinfletcher/apitest"
	"testing"
)

func GivenUserLoggedIn(t *testing.T) {
	accessToken := (*Login(t).Data()).AccessToken
	config.SetAccessToken(t, accessToken)
}

var IdentityHandler = apitest.New().Handler(pkg.Router())

func Login(t *testing.T) io.Response[*io.LoginResponse] {
	response := IdentityHandler.Post("/v1/login").Expect(t)
	result := response.End()
	return io.Response[*io.LoginResponse]{
		Response: response,
		Result:   result,
	}
}
