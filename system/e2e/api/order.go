package api

import (
	"github.com/steinfletcher/apitest"
	"order/pkg"
	"testing"
)

var OrderHandler = apitest.New().Handler(pkg.Router())

func CreateOrder(t *testing.T) *apitest.Response {
	return OrderHandler.Post("/v1/orders").Expect(t)
}

func CreateOrderReport(t *testing.T) *apitest.Response {
	return OrderHandler.Post("/v1/orders/report").Expect(t)
}
