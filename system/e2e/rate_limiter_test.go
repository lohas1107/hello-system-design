package e2e

import (
	"e2e/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
	"time"
)

type RateLimiterTestSuite struct {
	suite.Suite
}

func TestRateLimiterTestSuite(t *testing.T) {
	suite.Run(t, new(RateLimiterTestSuite))
}

func (s *RateLimiterTestSuite) Test_authenticated_api_request_count() {
	api.GivenUserLoggedIn(s.T())
	s.givenCreateOrderRequestCount(s.T(), 4)

	s.T().Run("below limit", func(t *testing.T) {
		actual := api.CreateOrder(t)
		s.statusShouldBe2XX(t, actual.Result.Response.StatusCode)
	})
	s.T().Run("exceeds limit", func(t *testing.T) {
		actual := api.CreateOrder(t)
		s.statusShouldBe429(t, actual.Result.Response.StatusCode)
	})
}

func (s *RateLimiterTestSuite) Test_unauthenticated_api_request_count() {
	api.GivenUserIp(s.T(), "127.0.0.1")
	s.givenLoginRequestCount(s.T(), 4)

	s.T().Run("below limit", func(t *testing.T) {
		actual := api.Login(t)
		s.statusShouldBe2XX(t, actual.Result.Response.StatusCode)
	})
	s.T().Run("exceeds limit", func(t *testing.T) {
		actual := api.Login(t)
		s.statusShouldBe429(t, actual.Result.Response.StatusCode)
	})
}

func (s *RateLimiterTestSuite) Test_custom_api_request_count() {
	path := "/v1/orders/reports"
	requestedAt := time.Now().Add(-time.Second).UTC().UnixMilli()
	api.GivenLastRequestedAt(path, requestedAt)

	s.T().Run("below limit", func(t *testing.T) {
		actual := api.CreateOrderReport(t)
		s.statusShouldBe2XX(t, actual.Result.Response.StatusCode)
	})
	s.T().Run("exceeds limit", func(t *testing.T) {
		actual := api.CreateOrderReport(t)
		s.statusShouldBe429(t, actual.Result.Response.StatusCode)
	})
}

func (s *RateLimiterTestSuite) givenCreateOrderRequestCount(t *testing.T, count int) {
	for range count {
		api.CreateOrder(t)
	}
}

func (s *RateLimiterTestSuite) givenLoginRequestCount(t *testing.T, count int) {
	for range count {
		api.Login(t)
	}
}

func (s *RateLimiterTestSuite) statusShouldBe2XX(t *testing.T, code int) {
	assert.LessOrEqual(t, 200, code)
	assert.Greater(t, 300, code)
}

func (s *RateLimiterTestSuite) statusShouldBe429(t *testing.T, code int) bool {
	return assert.Equal(t, http.StatusTooManyRequests, code)
}
