// Package error defined error struct.
//
// provides an error struct implementing the default error interface
// and constructors
package error

import (
	"fmt"
	"reflect"
)

// NewErrorf constructor
func NewErrorf(domain string, code string, msg string, a ...any) *Error {
	err := new(Error)
	err.Domain = domain
	err.Code = code
	if len(a) > 0 {
		fst := a[0]
		if !checkNilInterface(fst) {
			msg = fmt.Sprintf(msg, a...)
		}
	}

	err.Message = msg

	return err
}

func checkNilInterface(i interface{}) bool {
	iv := reflect.ValueOf(i)
	if !iv.IsValid() {
		return true
	}
	switch iv.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Func, reflect.Interface:
		return iv.IsNil()
	default:
		return false
	}
}

// NewErrorWithTracef constructor for eerors with trace data
func NewErrorWithTracef(domain string, code string, msg string, trace string, a ...any) *Error {
	err := NewErrorf(domain, code, msg, a...)
	err.Trace = trace
	return err
}
