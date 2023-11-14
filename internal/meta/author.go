// Package meta author file defines variables.
//
// provides variables to store author info and getters to give read access
package meta

import "github.com/cthiel77/server-demo/internal/meta/model"

var (
	authorName    = "Carsten Thiel"
	authorEmail   = "dev@thiel-inet.de"
	authorCompany = "yolo"
	//Author meta data
	Author = model.Author{
		Name:    GetAuthorName(),
		Email:   GetAuthorEmail(),
		Company: GetAuthorCompany(),
	}
)

// GetAuthorName returns author name
func GetAuthorName() string {
	return authorName
}

// GetAuthorEmail returns author email
func GetAuthorEmail() string {
	return authorEmail
}

// GetAuthorCompany returns author company
func GetAuthorCompany() string {
	return authorCompany
}
