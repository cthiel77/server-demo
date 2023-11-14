// Package meta version file defines variables.
//
// provides variables to store versin info and getters to give read access
package meta

import (
	"testing"
)

func TestGetVersionID(t *testing.T) {
	value := GetVersionID()
	expected := "undefined version id"
	if value != expected {
		t.Errorf("wrong value %s, expected %s", value, expected)
	}
}

func TestGetBuildTime(t *testing.T) {
	value := GetBuildTime()
	expected := "undefined buildtime"
	if value != expected {
		t.Errorf("wrong value %s, expected %s", value, expected)
	}
}

func TestGetVersionHashAlg(t *testing.T) {
	value := GetVersionHashAlg()
	expected := "undefined hash alg"
	if value != expected {
		t.Errorf("wrong value %s, expected %s", value, expected)
	}
}

func TestGetVersionHash(t *testing.T) {
	value := GetVersionHash()
	expected := "undefined hash"
	if value != expected {
		t.Errorf("wrong value %s, expected %s", value, expected)
	}
}
