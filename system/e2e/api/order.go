package api

import (
	"e2e/config"
	"github.com/steinfletcher/apitest"
	"order/pkg"
	"testing"
)

func CreateOrder(t *testing.T) *apitest.Response {
	return apitest.New().Debug().
		Handler(pkg.Router()).
		Post("/v1/orders").
		Headers(HttpHeaders(
			BearerToken(config.AccessToken()),
		)).
		Expect(t)
}

func CreateOrderReport(t *testing.T) *apitest.Response {
	return apitest.New().Debug().
		Handler(pkg.Router()).
		Post("/v1/orders/report").
		Header("Authentication", "Bearer 456").
		Expect(t)
}
