package cfg_test

import (
	"testing"

	"github.com/cthiel77/server-demo/internal/cfg"
	"github.com/stretchr/testify/assert"
)

func TestAppModeDefaultDev(t *testing.T) {

	mode := cfg.GetAppModeKey()

	assert.Equalf(
		t,
		cfg.ModeKeyDev,
		mode,
		"wrong mode %s. Expected mode dev", mode,
	)

}

func TestDevModeEnabled(t *testing.T) {

	tables := []struct {
		mode         string
		expectResult bool
	}{
		{cfg.ModeKeyDev, true},
		{cfg.ModeKeyProd, false},
		{"nonsense", false},
	}

	for _, row := range tables {
		cfg.Mode = row.mode
		res := cfg.DevModeEnabled()

		if row.expectResult {
			assert.Truef(t, res, "this should be dev mode: %s", row.mode)
		} else {
			assert.Falsef(t, res, "value must not be dev mode: %s", row.mode)
		}

	}
}

func TestProdModeEnabled(t *testing.T) {

	tables := []struct {
		mode         string
		expectResult bool
	}{
		{cfg.ModeKeyDev, false},
		{cfg.ModeKeyProd, true},
		{"nonsense", true},
	}

	for _, row := range tables {
		cfg.Mode = row.mode
		res := cfg.ProdModeEnabled()

		if row.expectResult {
			assert.Truef(t, res, "this should be prod mode: %s", row.mode)
		} else {
			assert.Falsef(t, res, "value must not be prod mode: %s", row.mode)
		}

	}
}
