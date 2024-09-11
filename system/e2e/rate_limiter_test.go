package e2e

import (
	"e2e/api"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type RateLimiterTestSuite struct {
	suite.Suite
}

func TestRateLimiterTestSuite(t *testing.T) {
	suite.Run(t, new(RateLimiterTestSuite))
}

func (s *RateLimiterTestSuite) Test_authenticated_api_request_count() {
	givenApiRequestCount(s.T(), api.CreateOrder, 4)

	s.T().Run("below limit", func(t *testing.T) {
		response := api.CreateOrder(s.T())
		s.statusShouldBe(response, http.StatusCreated)
	})
	s.T().Run("exceeds limit", func(t *testing.T) {
		response := api.CreateOrder(s.T())
		s.statusShouldBe(response, http.StatusTooManyRequests)
	})
}

func (s *RateLimiterTestSuite) Test_unauthenticated_api_request_count_below_limit() {
	for range 5 {
		response := api.Login(s.T())
		s.statusShouldBe(response, http.StatusOK)
	}
}

func (s *RateLimiterTestSuite) Test_custom_api_request_count_below_limit() {
	response := api.CreateOrderReport(s.T())
	s.statusShouldBe(response, http.StatusAccepted)
}

func givenApiRequestCount(t *testing.T, api func(t *testing.T) *apitest.Response, count int) {
	for range count {
		api(t).End()
	}
}

func (s *RateLimiterTestSuite) statusShouldBe(response *apitest.Response, status int) {
	response.Status(status).End()
}
