package pkg

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOrder(context *gin.Context) {
	context.Status(http.StatusCreated)
}

func CreateOrderReport(context *gin.Context) {
	context.Status(http.StatusAccepted)
}
