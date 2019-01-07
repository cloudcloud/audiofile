package data

import (
	"database/sql"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"

	// All migrations are pulled in as sql files
	_ "github.com/golang-migrate/migrate/v4/source/file"

	// Interfacing with the sqlite3 database
	_ "github.com/mattn/go-sqlite3"
)

func (d *Data) migrate() error {
	db, err := sql.Open("sqlite3", "./.audiofile.db")
	if err != nil {
		return err
	}

	driver, _ := sqlite3.WithInstance(db, &sqlite3.Config{})

	// TODO: use go-bindata for these migrations
	m, _ := migrate.NewWithDatabaseInstance(
		"file://data/migrations",
		"sqlite3",
		driver,
	)

	err = m.Up()
	if err != nil && err.Error() != "no change" {
		return err
	}

	return nil
}
