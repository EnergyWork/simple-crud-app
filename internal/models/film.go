package models

import (
	"time"

	errs "simple-crud-app/internal/lib/errors"
)

type Film struct {
	ID          uint64
	Type        *string
	Name        string
	ReleaseDate *time.Time
	Duration    *string
	SerialCount *uint64
	Score       *uint64
	Comment     *string
}

func (f *Film) Create(db DB) *errs.Error {
	sqlStr := `INSERT INTO film(type, name, year, duration, serial_count, score, comment) VALUES($1, $2, $3, $4, $5, $6, $7)`
	_, err := db.Exec(sqlStr, f.Type, f.Name, f.ReleaseDate, f.Duration, f.SerialCount, f.Score, f.Comment)
	if err != nil {
		return errs.New().SetCode(errs.ERROR_INTERNAL).SetMsg("failed to create a new record: %s", err)
	}
	return nil
}
