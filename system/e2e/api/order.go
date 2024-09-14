package api

import (
	"e2e/config"
	"gateway/pkg/io"
	"gateway/pkg/router"
	"github.com/steinfletcher/apitest"
	"testing"
)

func CreateOrder(t *testing.T) io.Response[*any] {
	result := apitest.New().
		Handler(router.Order()).
		Post("/v1/orders").
		Headers(io.HttpHeaders(
			io.BearerToken(config.AccessToken()),
		)).
		Expect(t).
		End()

	return io.Response[*any]{Result: result}
}

func CreateOrderReport(t *testing.T) io.Response[*any] {
	result := apitest.New().
		Handler(router.Order()).
		Post("/v1/orders/reports").
		Headers(io.HttpHeaders(
			io.BearerToken(config.AccessToken()),
		)).
		Expect(t).
		End()

	return io.Response[*any]{Result: result}
}
