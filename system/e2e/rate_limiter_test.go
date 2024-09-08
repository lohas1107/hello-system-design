package e2e

import (
	"github.com/gin-gonic/gin"
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

func (s *RateLimiterTestSuite) SetupTest() {

}

func Router() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("v1")
	v1.POST("/orders", CreateOrder)
	return router
}

func CreateOrder(context *gin.Context) {
	context.Status(http.StatusCreated)
}

func (s *RateLimiterTestSuite) Test_authenticated_api_request_count_below_limit() {
	apitest.New().Debug().
		Handler(Router()).
		Post("http://localhost:8080/v1/orders").
		Expect(s.T()).
		Status(http.StatusCreated).
		End()
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
