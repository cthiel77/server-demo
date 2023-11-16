package db

// HeroSet a country dataset to be stored
type HeroSet struct {
	ID        uint64 `json:"id" example:"123"`
	FirstName string `json:"firstName" example:"Wade"`
	LastName  string `json:"lastName" example:"Wilson"`
	HeroName  string `json:"heroName" example:"Deadpool"`
}
