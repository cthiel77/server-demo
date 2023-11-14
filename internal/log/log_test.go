package log

import (
	"testing"

	errors "github.com/cthiel77/server-demo/internal/error"
	"github.com/cthiel77/server-demo/internal/meta/model"
)

func TestGetLogLevel(t *testing.T) {
	InitFlags()

	lvl := LogLevel

	if lvl != "info" {
		t.Errorf("wrong level %s. Expected info", lvl)
		return
	}
	t.Logf("ok, got level %s", lvl)
}

func TestLoggerWarn(t *testing.T) {
	value := true

	l := NewLogger(model.AppMeta{Name: "testApp"})
	l.Errorf("test error")
	l.Warnf("test warning")
	l.Infof("test info")
	l.Debugf("test debug")
	l.Tracef("test trace")
	l.Printf("test print")

	if !value {
		t.Errorf("false")
	} else {
		t.Logf("done")
	}
}

func TestLoggerError(t *testing.T) {

	l := NewLogger(model.AppMeta{Name: "testApp"})

	// custom test case template
	// changed to fit your needs
	type errorTestCases struct {
		description string // add a description for the
		// case to have transparency
		input   any
		param   any
		msgType string
	}

	for _, scenario := range []errorTestCases{
		{
			description: "error default",
			input:       errors.NewError("1684872539", errors.StatusBadGateway, "this is an error"),
			msgType:     "error",
		},
		{
			description: "error string only",
			input:       "et sadipscing",
			msgType:     "error",
		},
		{
			description: "error string with param error",
			input:       "et sadipscing %s",
			param:       errors.NewError("1684872539", errors.StatusBadGateway, "this is an error"),
			msgType:     "errorf",
		},
		{
			description: "warn default",
			input:       errors.NewError("1684872539", errors.StatusBadGateway, "this is an error"),
			msgType:     "warn",
		},
		{
			description: "warn string only",
			input:       "et sadipscing",
			msgType:     "warn",
		},
		{
			description: "warn string with param error",
			input:       "et sadipscing %s",
			param:       errors.NewError("1684872539", errors.StatusBadGateway, "this is an error"),
			msgType:     "warnf",
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			// test the function

			switch scenario.msgType {
			case "error":
				l.Error(scenario.input)
			case "errorf":
				l.Errorf(scenario.input.(string), scenario.param)
			case "warn":
				l.Warn(scenario.input)
			case "warnf":
				l.Warnf(scenario.input.(string), scenario.param)
			}

		})
	}
}

func TestLoggerWarn2(t *testing.T) {
	value := true

	l := NewLogger(model.AppMeta{Name: "testApp"})
	l.Warn(errors.NewError("1684872539", errors.StatusBadGateway, "this is an error"))

	if !value {
		t.Errorf("false")
	} else {
		t.Logf("done")
	}
}

func TestLoggerUndefinedLevelDefaultInfo(t *testing.T) {

	l := NewLogger(model.AppMeta{Name: "testApp"})
	lvl := l.GetLevel()

	if lvl.String() != "info" {
		t.Errorf("wrong default level %s. Expected info", lvl.String())
		return
	}
}
