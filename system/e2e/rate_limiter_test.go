package e2e

import (
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/suite"
	identity "identity/pkg"
	"net/http"
	order "order/pkg"
	"testing"
)

type RateLimiterTestSuite struct {
	suite.Suite
}

func TestRateLimiterTestSuite(t *testing.T) {
	suite.Run(t, new(RateLimiterTestSuite))
}

func (s *RateLimiterTestSuite) SetupTest() {

}

func (s *RateLimiterTestSuite) Test_authenticated_api_request_count_below_limit() {
	apitest.New().
		Handler(order.Router()).
		Post("/v1/orders").
		Expect(s.T()).
		Status(http.StatusCreated).
		End()
}

func (s *RateLimiterTestSuite) Test_unauthenticated_api_request_count_below_limit() {
	apitest.New().
		Handler(identity.Router()).
		Post("/v1/login").
		Expect(s.T()).
		Status(http.StatusOK).
		End()
}

func (s *RateLimiterTestSuite) Test_api_request_count_below_limit() {
	s.T().Run("custom", func(t *testing.T) {

	})
}

func (s *RateLimiterTestSuite) Test_api_request_count_exceeds_limit() {
	s.T().Run("authenticated", func(t *testing.T) {

	})

	s.T().Run("unauthenticated", func(t *testing.T) {

	})

	s.T().Run("custom", func(t *testing.T) {

	})
}
