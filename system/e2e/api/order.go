package api

import (
	"github.com/steinfletcher/apitest"
	"order/pkg"
	"testing"
)

func CreateOrder(t *testing.T, headers map[string]string) *apitest.Response {
	return apitest.New().Debug().
		Handler(pkg.Router()).
		Post("/v1/orders").
		Headers(headers).
		Expect(t)
}

func CreateOrderReport(t *testing.T) *apitest.Response {
	return apitest.New().Debug().
		Handler(pkg.Router()).
		Post("/v1/orders/report").
		Header("Authentication", "Bearer 456").
		Expect(t)
}
