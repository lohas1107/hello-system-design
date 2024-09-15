package mock

import (
	"os"
	"testing"
)

func SetAccessToken(t *testing.T, value string) {
	t.Setenv("accessToken", value)
}

func AccessToken() string {
	return os.Getenv("accessToken")
}

func SetClientIp(t *testing.T, value string) {
	t.Setenv("clientIp", value)
}

func ClientIp() string {
	return os.Getenv("clientIp")
}
