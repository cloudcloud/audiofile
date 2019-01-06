package audiofile

import (
	"time"

	"github.com/golang-plus/uuid"
)

// Directory defines a location from which files will be discovered.
type Directory struct {
	DateAdded   time.Time `json:"date_added" db:"date_added"`
	DateUpdated time.Time `json:"date_updated" db:"date_updated"`
	Directory   string    `json:"directory"`
	ID          string    `json:"id"`
}

// GenerateID will provide a generated UUIDv4 value to the ID attribute if
// the attribute does not already have a value.
func (d Directory) GenerateID() (Directory, error) {
	if d.ID == "" || len(d.ID) < 1 {
		u, err := uuid.NewV4()
		if err == nil {
			d.ID = u.String()
		}
	}

	return d, nil
}
