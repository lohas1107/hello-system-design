package e2e

import (
	"github.com/gin-gonic/gin"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/suite"
	"net/http"
	"order/pkg"
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
		Handler(pkg.Router()).
		Post("/v1/orders").
		Expect(s.T()).
		Status(http.StatusCreated).
		End()
}

func (s *RateLimiterTestSuite) Test_unauthenticated_api_request_count_below_limit() {
	apitest.New().
		Handler(Router()).
		Post("/v1/login").
		Expect(s.T()).
		Status(http.StatusOK).
		End()
}

func Router() *gin.Engine {
	engine := gin.Default()
	v1 := engine.Group("v1")
	v1.POST("/login")
	return engine
}

func (s *RateLimiterTestSuite) Test_api_request_count_below_limit() {
	s.T().Run("unauthenticated", func(t *testing.T) {

	})

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
