package db

import (
	"github.com/cthiel77/server-demo/internal/db/sampledata"
	defaultError "github.com/cthiel77/server-demo/internal/error"
	"github.com/hashicorp/go-memdb"
)

func loadHeroData(s *memdb.MemDB) *defaultError.Error {

	//db transaction mode
	// Create a write transaction
	txn := s.Txn(true)
	defer txn.Abort()

	//insert model data
	ID := 0

	list, e := sampledata.GetHeroData()
	if e != nil {
		return e
	}

	for _, ci := range list {
		ID = ID + 1

		m := HeroSet{
			ID:        uint64(ID),
			FirstName: ci.FirstName,
			LastName:  ci.LastName,
			HeroName:  ci.HeroName,
		}
		if err := txn.Insert(TableNameHero, m); err != nil {
			return defaultError.FromGoError("1700066262", defaultError.StatusNotAcceptableData, err)
		}
	}
	// Commit the transaction
	txn.Commit()
	txn.Abort()

	return nil
}

///////////////////////////////////////
// Heroinfo
///////////////////////////////////////

type HeroQueryStack struct {
	Session *memdb.MemDB
}

// NewHeroQueryStack creates a new query stack
func NewHeroQuerstack(s *memdb.MemDB) HeroQueryStack {
	return HeroQueryStack{
		Session: s,
	}
}

// GetAll returns all country info
func (ci HeroQueryStack) GetAll() (list []HeroSet, err *defaultError.Error) {

	// Create read-only transaction
	txn := ci.Session.Txn(false)
	defer txn.Abort()

	rawRes, dbErr := txn.Get(TableNameHero, "heroName")
	if dbErr != nil {
		err = defaultError.FromGoError("1700066267", defaultError.StatusNotAcceptableData, dbErr)
		return
	}
	if rawRes == nil {
		err = defaultError.NewError("17000662720", defaultError.StatusNotFound, "no Heroinfo found")
		return
	}

	for obj := rawRes.Next(); obj != nil; obj = rawRes.Next() {
		c := obj.(HeroSet)
		list = append(list, c)
	}

	return
}

// GetByIDreturns a dataset or an error id not found
func (ci HeroQueryStack) GetByID(s uint64) (*HeroSet, *defaultError.Error) {

	// Create read-only transaction
	txn := ci.Session.Txn(false)
	defer txn.Abort()

	rawRes, err := txn.First(TableNameHero, "id", s)
	if err != nil {
		return nil, defaultError.FromGoError("1700066277", defaultError.StatusNotAcceptableData, err)
	}
	if rawRes == nil {
		return nil, defaultError.NewErrorf("1700066280", defaultError.StatusNotFound, "no data found for id %d", s)
	}

	ds := rawRes.(HeroSet)

	return &ds, nil
}

// Insert inserts a dataset or an error id not found
func (ci HeroQueryStack) Insert(id uint64, firstName, lastName, heroName string) *defaultError.Error {

	// Create read-only transaction
	txn := ci.Session.Txn(true)
	defer txn.Abort()

	if e := txn.Insert(TableNameHero, HeroSet{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		HeroName:  heroName,
	}); e != nil {
		return defaultError.FromGoError("1700067629", defaultError.StatusQueryWriteError, e)
	}

	// Commit the transaction
	txn.Commit()
	txn.Abort()

	return nil
}
