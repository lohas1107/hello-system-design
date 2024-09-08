package pkg

import (
	"github.com/gin-gonic/gin"
	"order/internal/router"
)

func Router() *gin.Engine {
	return router.Engine()
}
