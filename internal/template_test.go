package internal_test

import (
	"testing"

	"github.com/cthiel77/server-demo/internal"
	"github.com/stretchr/testify/assert"
)

func TestRenderTemplate(t *testing.T) {

	tplContent := `heallo this a sample content with variables like a={{.a}} and b={{.b}}`

	// custom test case template
	// changed to fit your needs
	type errorTestCases struct {
		description string // add a description for the
		// case to have transparency
		data      map[string]interface{}
		expectErr bool
	}

	for _, scenario := range []errorTestCases{
		{
			description: "default",
			data: map[string]interface{}{
				"a": 1,
				"b": 2,
			},
			expectErr: false,
		},
		{
			description: "missing variable",
			data: map[string]interface{}{
				"a": 1,
			},
			expectErr: true,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			s, err := internal.RenderTemplate(tplContent, scenario.data)
			if err != nil {
				if scenario.expectErr {
					t.Logf("success! expected this error: %+v", err)
				} else {
					t.Errorf("failure! %+v", err)
				}
				return
			}

			assert.Nil(t, err)
			assert.NotNil(t, s)

		})
	}
}
