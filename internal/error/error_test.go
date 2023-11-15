// Package error defined error struct.
//
// provides an error struct implementing the default error interface
// and constructors
package error_test

import (
	"regexp"
	"strconv"
	"testing"

	defaultError "github.com/cthiel77/server-demo/internal/error"
	"github.com/stretchr/testify/assert"
)

func getStandardError() error {
	_, e := strconv.Atoi("this wont do a conversion to int")
	return e
}

func TestNewError(t *testing.T) {

	err := defaultError.NewError("1699199734", defaultError.StatusNotFound, "new error")
	expectedCode := defaultError.StatusNotFound
	if err.GetCode() != expectedCode {
		t.Errorf("wrong code %s, expected %s", err.GetCode(), expectedCode)
	}

}

func TestNewErrorWithTrace(t *testing.T) {
	err := defaultError.NewErrorWithTrace("1699199739", defaultError.StatusNotFound, "new error", "trace")
	expectedTrace := "trace"
	if err.Trace != expectedTrace {
		t.Errorf("wrong code %q, expected %q", err.Trace, expectedTrace)
	}
}

func TestFromGoError(t *testing.T) {
	e := getStandardError()
	err := defaultError.FromGoError("1699199743", defaultError.StatusNotFound, e)
	expectedCode := defaultError.StatusNotFound
	if err.GetCode() != expectedCode {
		t.Errorf("wrong code %s, expected %s", err.GetCode(), expectedCode)
	}
}

func TestToString(t *testing.T) {
	err := defaultError.FromGoError("1699199747", defaultError.StatusNotFound, getStandardError())
	expected := "[ERROR 404"
	s := err.ToString()
	if m, _ := regexp.MatchString(`^\[ERROR\s404`, s); !m {
		t.Errorf("wrong code %q, expected %q", s, expected)
	}
}

func TestErrorNil(t *testing.T) {
	err := defaultError.FromGoError("1699199751", defaultError.StatusNotFound, nil)
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
		input       *[]defaultError.Error
		expectError bool
	}

	for _, scenario := range []errorTestCases{
		{
			description: "default",
			input: &[]defaultError.Error{
				{Domain: "1699199755", Code: "400", Message: "error 1"},
				{Domain: "1699199760", Code: "500", Message: "error 2"},
			},
			expectError: true,
		},
		{
			description: "empty list",
			input:       &[]defaultError.Error{},
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

			err := defaultError.FromErrorList("1699199764", "400", scenario.input)

			if scenario.expectError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}

			t.Logf("%+v", err)
		})
	}
}
