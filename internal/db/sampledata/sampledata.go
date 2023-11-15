package sampledata

import (
	//_ embed files
	_ "embed"
	"encoding/json"

	defaultError "github.com/cthiel77/server-demo/internal/error"
)

var (
	//go:embed dbinit_heroes.json
	heroSampleFile []byte
)

// HeroSample represents a sample import data format
type HeroSample struct {
	ID        uint64 `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	HeroName  string `json:"heroName,omitempty"`
}

// GetHeroData loads sample data to insert into db
func GetHeroData() (list []HeroSample, err *defaultError.Error) {

	if e := json.Unmarshal(heroSampleFile, &list); e != nil {
		err = defaultError.FromGoError("1700084638", defaultError.StatusDataJSONDecodeError, e)
	}

	return
}
