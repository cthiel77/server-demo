package db_test

import (
	"testing"

	"github.com/cthiel77/server-demo/internal/db"
	"github.com/cthiel77/server-demo/internal/db/sampledata"
	"github.com/stretchr/testify/assert"
)

func TestDBRun(t *testing.T) {

	if err := db.Run(); err != nil {
		t.Fatalf("%+v", err)
	}
	t.Log("db started")
}

func TestHeroDataQueryAll(t *testing.T) {

	if err := db.Run(); err != nil {
		t.Fatalf("%+v", err)
	}

	// Create read-only transaction
	txn := db.DBSession.Txn(false)
	defer txn.Abort()

	// Lookup by heroName
	rawRes, err := txn.Get(db.TableNameHero, "heroName")
	if err != nil {
		t.Fatal(err)
	}
	list := make([]db.HeroSet, 0)
	for obj := rawRes.Next(); obj != nil; obj = rawRes.Next() {
		ds := obj.(db.HeroSet)
		list = append(list, ds)
	}
	t.Logf("%+v", list)
}

func TestHeroDataQueryInsert(t *testing.T) {

	if err := db.Run(); err != nil {
		t.Fatalf("%+v", err)
	}

	querystack := db.NewHeroQuerstack(db.DBSession)

	// custom test case template
	// changed to fit your needs
	type errorTestCases struct {
		description string // add a description for the
		// case to have transparency
		data      sampledata.HeroSample
		expectErr bool
	}

	for _, scenario := range []errorTestCases{
		{
			description: "default",
			data: sampledata.HeroSample{
				ID:        8,
				FirstName: "Peter",
				LastName:  "Quill",
				HeroName:  "Starlord",
			},
			expectErr: false,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			err := querystack.Insert(
				scenario.data.ID,
				scenario.data.FirstName,
				scenario.data.LastName,
				scenario.data.HeroName,
			)
			if err != nil {
				if scenario.expectErr {
					t.Logf("success! expected this error: %+v", err)
				} else {
					t.Errorf("failure! %+v", err)
				}
				return
			}

			assert.Nil(t, err)

			list, _ := querystack.GetAll()

			t.Logf("list: %+v", list)
		})
	}

}

