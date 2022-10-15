package models

import (
	errs "simple-crud-app/internal/lib/errors"
	"time"
)

type Season struct {
	ID        uint64
	SerialID  uint64
	Number    uint64
	Series    map[string]string // {"1": "42m", "2": "46m"}
	CreatedAt time.Time
	UpdatedAt *time.Time
}

func (obj *Season) Create(db DB) *errs.Error {
	sqlStr := `INSERT INTO seasons(serial_id, number, series) VALUES ($1, $2, $3)`
	_, err := db.Exec(sqlStr)
	if err != nil {
		return errs.New().SetCode(errs.ErrorInternal)
	}
	return nil
}
