package models

import (
	"time"

	errs "simple-crud-app/internal/lib/errors"
)

type Film struct {
	ID          uint64
	Name        string
	ReleaseDate *time.Time
	Duration    *string
	Score       *uint64
	Comment     *string
}

func (f *Film) Create(db DB) *errs.Error {
	sqlStr := `INSERT INTO film(name, year, duration, serial_count, score, comment) VALUES($1, $2, $3, $4, $5)`
	_, err := db.Exec(sqlStr, f.Name, f.ReleaseDate, f.Duration, f.Score, f.Comment)
	if err != nil {
		return errs.New().SetCode(errs.ERROR_INTERNAL).SetMsg("failed to create a new record: %s", err)
	}
	return nil
}