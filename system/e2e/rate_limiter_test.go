package e2e

import (
	"e2e/api"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/suite"
	"net/http"
	"os"
	"testing"
)

type RateLimiterTestSuite struct {
	suite.Suite
}

func TestRateLimiterTestSuite(t *testing.T) {
	suite.Run(t, new(RateLimiterTestSuite))
}

type HttpHeaderOption func(map[string]string)

func HttpHeaders(options ...HttpHeaderOption) map[string]string {
	headers := map[string]string{}
	for _, option := range options {
		option(headers)
	}
	return headers
}

func BearerToken(token string) HttpHeaderOption {
	return func(headers map[string]string) {
		headers["Authentication"] = "Bearer " + token
	}
}

func (s *RateLimiterTestSuite) Test_authenticated_api_request_count() {
	s.T().Setenv("accessToken", "123")
	headers := HttpHeaders(BearerToken(os.Getenv("accessToken")))
	s.givenCreateOrderRequestCount(s.T(), headers, 4)

	s.T().Run("below limit", func(t *testing.T) {
		response := api.CreateOrder(s.T(), headers)
		s.statusShouldBe(response, http.StatusCreated)
	})
	s.T().Run("exceeds limit", func(t *testing.T) {
		response := api.CreateOrder(s.T(), headers)
		s.statusShouldBe(response, http.StatusTooManyRequests)
	})
}

func (s *RateLimiterTestSuite) givenCreateOrderRequestCount(t *testing.T, headers map[string]string, count int) {
	for range count {
		api.CreateOrder(t, headers).End()
	}
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

func (s *RateLimiterTestSuite) statusShouldBe(response *apitest.Response, status int) {
	response.Status(status).End()
}
