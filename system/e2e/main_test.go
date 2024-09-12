package e2e

import (
	"e2e/config"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	config.Load()
	code := m.Run()
	os.Exit(code)
}
