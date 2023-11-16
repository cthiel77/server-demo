package response

import (
	defaultError "github.com/cthiel77/server-demo/internal/error"
)

// EErrorDatarror represents the error response
type ErrorData struct {
	Code string `json:"code" example:"10001"`
	Msg  string `json:"msg" example:"something went wrong"`
}

// NewFromDefaultError construct by default error model
func NewFromDefaultError(err defaultError.Error) (e ErrorData) {
	e.Code = err.Code
	e.Msg = err.Message

	return
}

// NewFromGoError construct by go error interface model
func NewFromGoError(code string, err error) (e ErrorData) {
	e.Code = code
	e.Msg = err.Error()

	return
}
