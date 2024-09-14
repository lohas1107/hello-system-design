package api

import (
	"e2e/config"
	"github.com/steinfletcher/apitest"
	"order/pkg"
	"testing"
)

func CreateOrder(t *testing.T) Response[*any] {
	response := apitest.New().Debug().
		Handler(pkg.Router()).
		Post("/v1/orders").
		Headers(HttpHeaders(
			BearerToken(config.AccessToken()),
		)).
		Expect(t)
	result := response.End()
	return Response[*any]{
		Response: response,
		Result:   result,
	}
}

func CreateOrderReport(t *testing.T) *apitest.Response {
	return apitest.New().Debug().
		Handler(pkg.Router()).
		Post("/v1/orders/report").
		Header("Authentication", "Bearer 456").
		Expect(t)
}
