// Package meta version file defines variables.
//
// provides variables to store versin info and getters to give read access
package meta

import "github.com/cthiel77/server-demo/internal/meta/model"

var (
	versionID      = "undefined version id"
	versionHashAlg = "undefined hash alg"
	versionHash    = "undefined hash"
	buildTime      = "undefined buildtime"
	//Version version data
	Version = model.Version{
		ID:        GetVersionID(),
		HashAlg:   GetVersionHashAlg(),
		Hash:      GetVersionHash(),
		BuildTime: GetBuildTime(),
	}
)

// GetVersionID returns the version string
func GetVersionID() string {
	return versionID
}

// GetBuildTime returns the build time
func GetBuildTime() string {
	return buildTime
}

// GetVersionHashAlg returns the  ersion hash algorhythm key
func GetVersionHashAlg() string {
	return versionHashAlg
}

// GetVersionHash returns the version hash
func GetVersionHash() string {
	return versionHash
}
