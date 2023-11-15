package model

//Version version object
type Version struct {
	ID        string // version
	HashAlg   string // hash revision used to build the program
	Hash      string // hash
	BuildTime string // when was the executable built
}
