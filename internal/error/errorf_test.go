// Package error defined error struct.
//
// provides an error struct implementing the default error interface
// and constructors
package error_test

import (
	"testing"

	omegaError "github.com/cthiel77/server-demo/internal/error"
	"github.com/stretchr/testify/assert"
)

func TestFromErrorf(t *testing.T) {
	// custom test case template
	// changed to fit your needs
	type errorTestCases struct {
		description string // add a description for the
		// case to have transparency
		msg          string
		arg1         any
		arg2         any
		arg3         any
		expectResult string
		expectError  bool
	}

	for _, scenario := range []errorTestCases{
		{
			description:  "default",
			msg:          "error: %s",
			arg1:         "yes",
			expectResult: "error: yes",
			expectError:  true,
		},
		{
			description:  "no args",
			msg:          "error: empty",
			expectResult: "error: empty",
			expectError:  true,
		},
		{
			description:  "slice of strings args",
			msg:          "error: %s %s",
			arg1:         "hello",
			arg2:         "world",
			expectResult: "error: hello world",
			expectError:  true,
		},
		{
			description:  "mixed args",
			msg:          "error: %d %s",
			arg1:         123,
			arg2:         "world",
			expectResult: "error: 123 world",
			expectError:  true,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			// test the function
			var err *omegaError.Error

			switch {
			case (scenario.arg3 != nil && scenario.arg2 != nil && scenario.arg1 != nil):

				err = omegaError.NewErrorf(
					"1684146751",
					omegaError.StatusNotAcceptable,
					scenario.msg,
					scenario.arg1,
					scenario.arg2,
					scenario.arg3,
				)
			case (scenario.arg2 != nil && scenario.arg1 != nil):
				err = omegaError.NewErrorf(
					"1684146751",
					omegaError.StatusNotAcceptable,
					scenario.msg,
					scenario.arg1,
					scenario.arg2,
				)
			default:
				err = omegaError.NewErrorf(
					"1684146751",
					omegaError.StatusNotAcceptable,
					scenario.msg,
					scenario.arg1,
				)

			}
			assert.NotNilf(t, err, err.Message)

			assert.Equal(t, scenario.expectResult, err.Message)

			assert.NotEmpty(t, err.GetCode())
			assert.NotEmpty(t, err.GetMessage())

		})
	}
}

func TestFromErrorWitrhTracef(t *testing.T) {
	// custom test case template
	// changed to fit your needs
	type errorTestCases struct {
		description string // add a description for the
		// case to have transparency
		msg          string
		trace        string
		arg1         any
		arg2         any
		arg3         any
		expectResult string
		expectError  bool
	}

	for _, scenario := range []errorTestCases{
		{
			description:  "default",
			msg:          "error: %s",
			trace:        "this is a trace",
			arg1:         "yes",
			expectResult: "error: yes",
			expectError:  true,
		},
		{
			description:  "no args",
			msg:          "error: empty",
			trace:        "this is a trace",
			expectResult: "error: empty",
			expectError:  true,
		},
		{
			description:  "slice of strings args",
			msg:          "error: %s %s",
			trace:        "this is a trace",
			arg1:         "hello",
			arg2:         "world",
			expectResult: "error: hello world",
			expectError:  true,
		},
		{
			description:  "mixed args",
			msg:          "error: %d %s",
			trace:        "this is a trace",
			arg1:         123,
			arg2:         "world",
			expectResult: "error: 123 world",
			expectError:  true,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			// test the function
			var err *omegaError.Error

			switch {
			case (scenario.arg3 != nil && scenario.arg2 != nil && scenario.arg1 != nil):

				err = omegaError.NewErrorWithTracef(
					"1684146751",
					omegaError.StatusNotAcceptable,
					scenario.msg,
					scenario.trace,
					scenario.arg1,
					scenario.arg2,
					scenario.arg3,
				)
			case (scenario.arg2 != nil && scenario.arg1 != nil):
				err = omegaError.NewErrorWithTracef(
					"1684146751",
					omegaError.StatusNotAcceptable,
					scenario.msg,
					scenario.trace,
					scenario.arg1,
					scenario.arg2,
				)
			default:
				err = omegaError.NewErrorWithTracef(
					"1684146751",
					omegaError.StatusNotAcceptable,
					scenario.msg,
					scenario.trace,
					scenario.arg1,
				)

			}
			assert.NotNilf(t, err, err.Message)

			assert.Equal(t, scenario.expectResult, err.Message)

			assert.NotEmpty(t, err.GetCode())
			assert.NotEmpty(t, err.GetMessage())

		})
	}
}
