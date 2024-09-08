package api

import (
	"github.com/steinfletcher/apitest"
	"identity/pkg"
	"testing"
)

var IdentityHandler = apitest.New().Handler(pkg.Router())

func Login(t *testing.T) *apitest.Response {
	return IdentityHandler.Post("/v1/login").Expect(t)
}
