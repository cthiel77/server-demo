// Package meta author file defines variables.
//
// provides variables to store author info and getters to give read access
package meta_test

import (
	"testing"

	"github.com/cthiel77/server-demo/internal/meta"
)

func TestGetAuthorNameInit(t *testing.T) {
	value := meta.GetAuthorName()
	expected := meta.Author.Name
	if value != expected {
		t.Errorf("wrong value %s, expected %s", value, expected)
	}
}

func TestGetAuthorEmail(t *testing.T) {
	value := meta.GetAuthorEmail()
	expected := meta.Author.Email
	if value != expected {
		t.Errorf("wrong value %s, expected %s", value, expected)
	}
}

func TestGetAuthorCompany(t *testing.T) {
	value := meta.GetAuthorCompany()
	expected := meta.Author.Company
	if value != expected {
		t.Errorf("wrong value %s, expected %s", value, expected)
	}
}
