package api

import (
	"encoding/json"
	"github.com/steinfletcher/apitest"
	"identity/pkg"
	"io"
	"testing"
)

var IdentityHandler = apitest.New().Handler(pkg.Router())

func Login(t *testing.T) Response[*LoginResponse] {
	response := IdentityHandler.Post("/v1/login").Expect(t)
	result := response.End()
	return Response[*LoginResponse]{
		Response: response,
		Result:   result,
	}
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}

type Response[T any] struct {
	Response *apitest.Response
	Result   apitest.Result
}

func (r Response[T]) Data() (data T) {
	body, _ := io.ReadAll(r.Result.Response.Body)
	_ = json.Unmarshal(body, &data)
	return
}
