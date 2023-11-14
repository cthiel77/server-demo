// Package error defined error struct.
//
// provides an error struct implementing the default error interface
// and constructors
package error_test

import (
	"regexp"
	"strconv"
	"testing"

	omegaError "github.com/cthiel77/server-demo/internal/error"
	"github.com/stretchr/testify/assert"
)

func getStandardError() error {
	_, e := strconv.Atoi("this wont do a conversion to int")
	return e
}

func TestNewError(t *testing.T) {

	err := omegaError.NewError("1699199734", "404", "new error")
	expectedCode := "404"
	if err.GetCode() != expectedCode {
		t.Errorf("wrong code %s, expected %s", err.GetCode(), expectedCode)
	}

}

func TestNewErrorWithTrace(t *testing.T) {
	err := omegaError.NewErrorWithTrace("1699199739", "404", "new error", "trace")
	expectedTrace := "trace"
	if err.Trace != expectedTrace {
		t.Errorf("wrong code %q, expected %q", err.Trace, expectedTrace)
	}
}

func TestFromGoError(t *testing.T) {
	e := getStandardError()
	err := omegaError.FromGoError("1699199743", "404", e)
	expectedCode := "404"
	if err.GetCode() != expectedCode {
		t.Errorf("wrong code %s, expected %s", err.GetCode(), expectedCode)
	}
}

func TestToString(t *testing.T) {
	err := omegaError.FromGoError("1699199747", "404", getStandardError())
	expected := "[ERROR 404"
	s := err.ToString()
	if m, _ := regexp.MatchString(`^\[ERROR\s404`, s); !m {
		t.Errorf("wrong code %q, expected %q", s, expected)
	}
}

func TestErrorNil(t *testing.T) {
	err := omegaError.FromGoError("1699199751", "404", nil)
	expected := "nil error given"
	s := err.ToString()
	assert.Contains(t, s, expected)
}

func TestFromErrorList(t *testing.T) {
	// custom test case template
	// changed to fit your needs
	type errorTestCases struct {
		description string // add a description for the
		// case to have transparency
		input       *[]omegaError.Error
		expectError bool
	}

	for _, scenario := range []errorTestCases{
		{
			description: "default",
			input: &[]omegaError.Error{
				{Domain: "1699199755", Code: "400", Message: "error 1"},
				{Domain: "1699199760", Code: "500", Message: "error 2"},
			},
			expectError: true,
		},
		{
			description: "empty list",
			input:       &[]omegaError.Error{},
			expectError: false,
		},
		{
			description: "no error list",
			input:       nil,
			expectError: false,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			// test the function

			err := omegaError.FromErrorList("1699199764", "400", scenario.input)

			if scenario.expectError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}

			t.Logf("%+v", err)
		})
	}
}
