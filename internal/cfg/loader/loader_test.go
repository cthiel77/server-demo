package loader_test

import (
	"testing"

	"github.com/cthiel77/server-demo/internal/cfg/loader"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfigDataFile(t *testing.T) {
	// custom test case template
	// changed to fit your needs
	type errorTestCases struct {
		description string // add a description for the
		// case to have transparency
		input       string
		expectError bool
	}

	for _, scenario := range []errorTestCases{
		// test case 1
		{
			description: "file does not exist",
			input:       "no_file",
			expectError: true,
		},
		{
			description: "file invalid",
			input:       "../../go.mod",
			expectError: true,
		},
		{
			description: "file valid",
			input:       "loader_test.json",
			expectError: false,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			// test the function
			_, msg, err := loader.LoadConfigDataFile(&scenario.input)

			if scenario.expectError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.NotEmpty(t, msg)
			}
		})
	}
}
