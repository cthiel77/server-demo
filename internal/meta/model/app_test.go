package model_test

import (
	"testing"

	"github.com/cthiel77/server-demo/internal/meta/model"
)

func TestApp(t *testing.T) {
	a := model.AppMeta{
		Name: "new",
	}
	expectName := "new"
	if a.Name != expectName {
		t.Errorf("wrong value %s, expected %s", a.Name, expectName)
	}
}

// Extended extended config data
func TestExtended(t *testing.T) {
	a := model.Extended{
		License: "new",
	}
	expectLicense := "new"
	if a.License != expectLicense {
		t.Errorf("wrong value %s, expected %s", a.License, expectLicense)
	}
}
