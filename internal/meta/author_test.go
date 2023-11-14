// Package meta author file defines variables.
//
// provides variables to store author info and getters to give read access
package meta

import (
	"testing"
)

func TestGetAuthorNameInit(t *testing.T) {
	value := GetAuthorName()
	expected := "Carsten Thiel"
	if value != expected {
		t.Errorf("wrong value %s, expected %s", value, expected)
	}
}

func TestGetAuthorEmail(t *testing.T) {
	value := GetAuthorEmail()
	expected := "dev@thiel-inet.de"
	if value != expected {
		t.Errorf("wrong value %s, expected %s", value, expected)
	}
}

func TestGetAuthorCompany(t *testing.T) {
	value := GetAuthorCompany()
	expected := "yolo"
	if value != expected {
		t.Errorf("wrong value %s, expected %s", value, expected)
	}
}
