package pkg

import (
	"github.com/gin-gonic/gin"
	"identity/internal/router"
)

func Router() *gin.Engine {
	return router.Engine()
}
