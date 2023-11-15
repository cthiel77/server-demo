package flags_test

import (
	"testing"

	"github.com/cthiel77/server-demo/internal/cfg"
	"github.com/cthiel77/server-demo/internal/cli/flags"
	"github.com/cthiel77/server-demo/internal/log"
)

func TestInit(t *testing.T) {
	flags.Init()

	expectMode := "dev"
	if cfg.Mode != expectMode {
		t.Errorf("wrong value %s, expected %s", cfg.Mode, expectMode)
	}

	expectLogLevel := "info"
	if log.LogLevel != expectLogLevel {
		t.Errorf("wrong value %s, expected %s", log.LogLevel, expectLogLevel)
	}

}

func TestReInit(t *testing.T) {
	flags.Init()
	flags.Init()

}
