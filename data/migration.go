package data

import (
	"database/sql"
	"log"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"

	// Interfacing with the sqlite3 database
	_ "github.com/mattn/go-sqlite3"
)

func (d *Data) migrate() error {
	db, err := sql.Open("sqlite3", "./.audiofile.db")
	if err != nil {
		log.Fatalf("Couldn't open DB: %s", err)
	}

	source, err := bindata.WithInstance(bindata.Resource(
		AssetNames(),
		func(name string) ([]byte, error) {
			return Asset(name)
		},
	))
	if err != nil {
		log.Fatalf("Unable to work with bindata: %s", err)
	}

	data, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatalf("Couldn't connect: %s", err)
	}

	m, err := migrate.NewWithInstance(
		"go-bindata",
		source,
		"sqlite3",
		data,
	)
	if err != nil {
		log.Fatalf("Unable to migrate: %s", err)
	}

	err = m.Up()
	if err != nil && err.Error() != "no change" {
		return err
	}

	return nil
}
