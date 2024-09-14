package api

import (
	"e2e/config"
	"gateway/pkg/io"
	"gateway/pkg/router"
	"github.com/steinfletcher/apitest"
	"testing"
)

func CreateOrder(t *testing.T) io.Response[*any] {
	response := apitest.New().
		Handler(router.Order()).
		Post("/v1/orders").
		Headers(io.HttpHeaders(
			io.BearerToken(config.AccessToken()),
		)).
		Expect(t)

	return io.Response[*any]{
		Response: response,
		Result:   response.End(),
	}
}

func CreateOrderReport(t *testing.T) io.Response[*any] {
	response := apitest.New().
		Handler(router.Order()).
		Post("/v1/orders/report").
		Headers(io.HttpHeaders(
			io.BearerToken(config.AccessToken()),
		)).
		Expect(t)

	return io.Response[*any]{
		Response: response,
		Result:   response.End(),
	}
}
