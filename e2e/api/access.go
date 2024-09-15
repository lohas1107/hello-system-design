package api

import (
	"e2e/mock"
	"gateway/pkg/io"
	"gateway/pkg/router"
	"github.com/steinfletcher/apitest"
	"testing"
)

func Login[T *io.LoginResponse](t *testing.T) io.Response[T] {
	result := apitest.New().
		Handler(router.Access()).
		Post("/v1/login").
		Headers(io.HttpHeaders(
			io.ClientIp(mock.ClientIp()),
		)).
		Expect(t).
		End()

	return io.Response[T]{Result: result}
}
