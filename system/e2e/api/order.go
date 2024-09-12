package api

import (
	"github.com/steinfletcher/apitest"
	"order/pkg"
	"os"
	"testing"
)

func CreateOrder(t *testing.T) *apitest.Response {
	return apitest.New().Debug().
		Handler(pkg.Router()).
		Post("/v1/orders").
		Headers(HttpHeaders(
			BearerToken(os.Getenv("accessToken")),
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
