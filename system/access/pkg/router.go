package pkg

import (
	"access/internal/router"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	return router.Engine()
}
