package db

import (
	defaultError "github.com/cthiel77/server-demo/internal/error"
	"github.com/hashicorp/go-memdb"
	log "github.com/sirupsen/logrus"
)

var (
	//DBSession the memory db session
	DBSession *memdb.MemDB
)

// table names
const (
	TableNameHero = "heroes"
)

func initDB() (*memdb.MemDB, *defaultError.Error) {

	tblSchema := make(map[string]*memdb.TableSchema)

	tblSchema[TableNameHero] = &memdb.TableSchema{
		Name: TableNameHero,
		Indexes: map[string]*memdb.IndexSchema{
			"id": {
				Name:    "id",
				Unique:  true,
				Indexer: &memdb.UintFieldIndex{Field: "ID"},
			},
			"heroName": {
				Name:    "heroName",
				Unique:  true,
				Indexer: &memdb.StringFieldIndex{Field: "HeroName"},
			},
		},
	}

	// Create the DB schema
	dbSchema := &memdb.DBSchema{
		Tables: tblSchema,
	}

	// Create a new data base
	db, err := memdb.NewMemDB(dbSchema)
	if err != nil {
		return nil, defaultError.FromGoError("1700066255", defaultError.StatusNotAcceptableData, err)
	}

	return db, nil
}

// Run startup
func Run() *defaultError.Error {

	if DBSession == nil {
		db, initErr := initDB()
		if initErr != nil {
			return initErr
		}
		DBSession = db
	}

	log.Info("loading hero data")
	if err := loadHeroData(DBSession); err != nil {
		return err
	}

	return nil
}
