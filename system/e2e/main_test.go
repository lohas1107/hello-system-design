package e2e

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	Load()
	code := m.Run()
	os.Exit(code)
}

func Load() {
	godotenv.Load()
	gin.SetMode(os.Getenv("GIN_MODE"))
}
