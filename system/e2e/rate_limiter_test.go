package e2e

import (
	"e2e/api"
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

func (s *RateLimiterTestSuite) SetupTest() {
}

func (s *RateLimiterTestSuite) Test_authenticated_api_request_count_below_limit() {
	response := api.CreateOrder(s.T())
	response.Status(http.StatusCreated).End()
}

func (s *RateLimiterTestSuite) Test_unauthenticated_api_request_count_below_limit() {
	response := api.Login(s.T())
	response.Status(http.StatusOK).End()
}

func (s *RateLimiterTestSuite) Test_custom_api_request_count_below_limit() {
	response := api.CreateOrderReport(s.T())
	response.Status(http.StatusAccepted).End()
}

func (s *RateLimiterTestSuite) Test_api_request_count_exceeds_limit() {
	s.T().Run("authenticated", func(t *testing.T) {

	})

	s.T().Run("unauthenticated", func(t *testing.T) {

	})

	s.T().Run("custom", func(t *testing.T) {

	})
}
