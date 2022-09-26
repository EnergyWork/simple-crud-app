package models

import (
	errs "simple-crud-app/internal/lib/errors"
	"time"
)

type Serial struct {
	ID          uint64
	UserID      uint64
	Name        string
	ReleaseDate *time.Time
	Score       *uint64
	Comment     *string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

func (f *Serial) Create(db DB) *errs.Error {
	sqlStr := `INSERT INTO serials (user_id, name, release_date, score, comment) VALUES($1, $2, $3, $4, $5, $6);`
	_, err := db.Exec(sqlStr, f.UserID, f.Name, f.ReleaseDate, f.Score, f.Comment)
	if err != nil {
		return errs.New().SetCode(errs.ERROR_INTERNAL).SetMsg("failed to create a new record: %s", err)
	}
	return nil
}
