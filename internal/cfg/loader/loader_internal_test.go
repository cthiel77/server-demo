package loader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestObj struct {
	Foo string
	O2  TestObj2
}

type TestObj2 struct {
	User string
	Pass string
	O3   TestObj3
}

type TestObj3 struct {
	Host string
	Port int
}

func TestPrintConfig(t *testing.T) {

	obj := TestObj{
		Foo: "bar",
		O2: TestObj2{
			User: "Max",
			Pass: "SecretPass",
			O3: TestObj3{
				Host: "myhost",
				Port: 123,
			},
		},
	}

	PrintConfig(obj)
}

func TestPrintConfigEmptyPass(t *testing.T) {

	obj := TestObj{
		Foo: "bar",
		O2: TestObj2{
			User: "Max",
			Pass: "",
			O3: TestObj3{
				Host: "myhost",
				Port: 123,
			},
		},
	}

	PrintConfig(obj)
}

func TestLoadConfigData(t *testing.T) {
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
			flagConfigFileName = scenario.input
			_, err := LoadConfigData()
			if scenario.expectError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
