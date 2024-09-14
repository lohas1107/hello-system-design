package pkg

import (
	"encoding/json"
	"github.com/steinfletcher/apitest"
	"io"
)

type Response[T any] struct {
	Response *apitest.Response
	Result   apitest.Result
}

func (r Response[T]) Data() (data T) {
	body, _ := io.ReadAll(r.Result.Response.Body)
	_ = json.Unmarshal(body, &data)
	return
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}
