package app

import (
	"testing"

	"github.com/cthiel77/server-demo/internal/cfg"
)

func TestAppModeDefaultDev(t *testing.T) {
	InitFlags()

	mode := cfg.Mode

	if mode != "dev" {
		t.Errorf("wrong mode %s. Expected mode dev", mode)
		return
	}
	t.Logf("ok, got mode %s", mode)

}

func TestPrintConfig(t *testing.T) {
	PrintConfig()
}
