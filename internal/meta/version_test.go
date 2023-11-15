// Package meta version file defines variables.
//
// provides variables to store versin info and getters to give read access
package meta_test

import (
	"testing"

	"github.com/cthiel77/server-demo/internal/meta"
)

func TestGetVersionID(t *testing.T) {
	value := meta.GetVersionID()
	expected := "undefined version id"
	if value != expected {
		t.Errorf("wrong value %s, expected %s", value, expected)
	}
}

func TestGetBuildTime(t *testing.T) {
	value := meta.GetBuildTime()
	expected := "undefined buildtime"
	if value != expected {
		t.Errorf("wrong value %s, expected %s", value, expected)
	}
}

func TestGetVersionHashAlg(t *testing.T) {
	value := meta.GetVersionHashAlg()
	expected := "undefined hash alg"
	if value != expected {
		t.Errorf("wrong value %s, expected %s", value, expected)
	}
}

func TestGetVersionHash(t *testing.T) {
	value := meta.GetVersionHash()
	expected := "undefined hash"
	if value != expected {
		t.Errorf("wrong value %s, expected %s", value, expected)
	}
}
