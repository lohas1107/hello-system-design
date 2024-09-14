package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
	"runtime"
)

func Load() {
	_, f, _, _ := runtime.Caller(0)
	path := filepath.Join(filepath.Dir(f), "./.env")
	fmt.Println("Loading .env file from", path)

	_ = godotenv.Load(path)
	gin.SetMode(os.Getenv("GIN_MODE"))
}
