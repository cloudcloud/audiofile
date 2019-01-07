// Package data provides abstraction for interacting with the
// data store of choice. In the case of audiofile, this will be
// with the local sqlite db.
package data

import (
	"fmt"

	"github.com/cloudcloud/audiofile"
	"github.com/jmoiron/sqlx"

	// Using the sqlite3 driver for sqlx
	_ "github.com/mattn/go-sqlite3"
)

// Data provides a struct to house the database connection so
// methods can be attached to work with data.
type Data struct {
	Errors []error
	conn   *sqlx.DB
}

// Open will setup the sqlite file and db handle.
func Open() (*Data, error) {
	db, err := sqlx.Open("sqlite3", "./.audiofile.db")
	if err != nil {
		return nil, err
	}

	d := &Data{conn: db}
	err = d.migrate()
	return d, err
}

// CaptureFile will take a full found file and capture the
// data points within the database.
func (d *Data) CaptureFile() error {
	return nil
}

// DeleteDirectory will accept an existing directory reference and remove
// it from the database.
func (d *Data) DeleteDirectory(dir audiofile.Directory) error {
	rows, err := d.conn.NamedExec(
		"delete from directories where id=:id",
		dir,
	)
	if err != nil {
		return err
	}
	r, err := rows.RowsAffected()
	if r != 1 {
		err = fmt.Errorf("multiple deletions for a single directory")
	}

	return err
}

// GetAlbums will retrieve a slice of Album entries.
func (d *Data) GetAlbums() []audiofile.Album {
	a := []audiofile.Album{}
	d.conn.Select(&a, "select * from albums")

	return a
}

// GetArtists will retrieve a slice of Artist entries.
func (d *Data) GetArtists() []audiofile.Artist {
	a := []audiofile.Artist{}
	d.conn.Select(&a, "select * from artists")

	return a
}

// GetDirectories will provide a list of all known directories.
func (d *Data) GetDirectories() ([]audiofile.Directory, error) {
	a := []audiofile.Directory{}
	d.conn.Select(&a, "select * from directories")

	return a, nil
}

// StoreDirectory will take a Directory and push it into the database.
func (d *Data) StoreDirectory(dir audiofile.Directory) (audiofile.Directory, error) {
	rows, err := d.conn.NamedExec(
		"replace into directories (id, directory, date_added, date_updated) values (:id, :directory, :date_added, :date_updated)",
		dir,
	)

	if err != nil {
		return dir, err
	}
	r, err := rows.RowsAffected()
	if err != nil {
		return dir, err
	}
	if r != 1 {
		return dir, fmt.Errorf("replace modified multiple rows")
	}

	return dir, nil
}
