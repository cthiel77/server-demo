// Package error defined error struct.
//
// provides an error struct implementing the default error interface
// and constructors
package error

import (
	"fmt"
	"strings"
)

// Error represents the rest-api error
type Error struct {
	Domain  string `json:"domain" example:"20230907120322"`
	Code    string `json:"code" example:"406000"`
	Message string `json:"message" example:"this is an error message!"`
	Trace   string `json:"trace"`
}

// Error error message implementation
func (e *Error) Error() string {
	if e == nil {
		return "no error message available"
	}
	return e.Message
}

// GetCode error code
func (e *Error) GetCode() string {
	if e == nil {
		return StatusUnexpectedError
	}
	return e.Code
}

// GetMessage error message
func (e *Error) GetMessage() string {
	if e == nil {
		return "no error message available"
	}
	return e.Message
}

// NewError constructor
func NewError(domain string, code string, msg string) *Error {
	err := new(Error)
	err.Domain = domain
	err.Code = code
	err.Message = fmt.Sprintf("Domain: %s Code: %s Message: %s", domain, code, msg)
	return err
}

// NewErrorWithTrace constructor for eerors with trace data
func NewErrorWithTrace(domain string, code string, msg string, trace string) *Error {
	err := NewError(domain, code, msg)
	err.Trace = trace
	return err
}

// FromGoError default error converter
func FromGoError(domain string, code string, err error) *Error {
	if err == nil {
		return NewError(domain, code, "nil error given")
	}
	return NewError(domain, code, err.Error())
}

// FromErrorList default error from error list
func FromErrorList(domain string, code string, list *[]Error) *Error {

	if list == nil {
		return nil
	}
	eList := []string{}
	for _, e := range *list {
		eList = append(eList, e.Message)
	}

	if len(eList) == 0 {
		return nil
	}

	return NewError(domain, code, strings.Join(eList, " | "))
}

// ToString convert error to string
func (e *Error) ToString() string {
	return fmt.Sprintf("[ERROR %s] - %s", e.Code, e.Error())
}
