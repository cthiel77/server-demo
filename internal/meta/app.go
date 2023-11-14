// Package meta app file defines variables.
//
// provides variables to store application info and getters to give read access
package meta

import (
	"encoding/json"
	"fmt"

	defaultError "github.com/cthiel77/server-demo/internal/error"
	"github.com/cthiel77/server-demo/internal/meta/model"
	catalogModel "github.com/cthiel77/server-demo/internal/meta/model/catalog"
	flag "github.com/spf13/pflag"
)

var (
	flagPrintLicense       bool
	flagPrintAuthorName    bool
	flagPrintAuthorCompany bool
	flagPrintVersion       bool

	appName        = "server demo"
	appDescription = "a server application demo including embedded sources"
	//App meta data
	App = model.AppMeta{
		Name:        GetAppName(),
		Description: GetAppDescription(),
	}

	extLicense = "undefined license"
	//Extended meta data
	Extended = model.Extended{
		License: GetAppLicense(),
	}
)

// InitAppCatalog init app catalog
func InitAppCatalog(appCatalogJSON string, catalog *catalogModel.App) *defaultError.Error {
	b := []byte(appCatalogJSON)
	convErr := json.Unmarshal(b, &catalog)
	if convErr != nil {
		return defaultError.FromGoError("1699200272", defaultError.StatusDataJSONDecodeError, convErr)
	}

	return nil
}

// InitFlags initialize flags
func InitFlags() {
	//define flags
	flag.BoolVar(&flagPrintAuthorName, "author", false, "")
	flag.BoolVar(&flagPrintAuthorCompany, "company", false, "")
	flag.BoolVar(&flagPrintLicense, "license", false, "")
	flag.BoolVarP(&flagPrintVersion, "version", "v", false, "")
}

// ProcessFlags process flags
func ProcessFlags() (processed bool) {
	if flagPrintLicense {
		fmt.Printf("%s\n", GetAppLicense())
		processed = true
	}
	if flagPrintAuthorName {
		fmt.Printf("%s<%s>\n", GetAuthorName(), GetAuthorEmail())
		processed = true
	}
	if flagPrintAuthorCompany {
		fmt.Printf("%s\n", GetAuthorCompany())
		processed = true
	}
	if flagPrintVersion {
		fmt.Printf("%s (%s)\n", GetVersionID(), GetBuildTime())
		processed = true
	}

	return
}

// GetAppName returns app name
func GetAppName() string {
	return appName
}

// GetAppDescription returns app name
func GetAppDescription() string {
	return appDescription
}

// GetAppLicense return license string
func GetAppLicense() string {
	return extLicense
}
