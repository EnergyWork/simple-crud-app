package models

import (
	"database/sql"
	"simple-crud-app/internal/domain"
	"time"

	errs "simple-crud-app/internal/lib/errors"
)

type Film struct {
	ID          uint64
	UserID      uint64
	Name        string
	ReleaseDate *time.Time
	Duration    *string
	Score       *uint64
	Comment     *string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	Watched     bool
}

func (f *Film) Create(db DB) *errs.Error {
	sqlStr := `INSERT INTO films (user_id, name, release_date, duration, score, comment, watched) VALUES($1, $2, $3, $4, $5, $6, $7);`
	_, err := db.Exec(sqlStr, f.UserID, f.Name, f.ReleaseDate, f.Duration, f.Score, f.Comment, f.Watched)
	if err != nil {
		return errs.New().SetCode(errs.ErrorInternal).SetMsg("failed to create a new record: %s", err)
	}
	return nil
}

func (f *Film) Update(db DB) *errs.Error {
	const sqlStr = `UPDATE films SET name=$1, release_date=$2, duration=$3, score=$4, comment=$5, updated_at=$7, watched=$8 WHERE id=$6;`
	_, err := db.Exec(sqlStr, f.Name, f.ReleaseDate, f.Duration, f.Score, f.Comment, f.ID, time.Now(), f.Watched)
	if err != nil {
		if err == sql.ErrNoRows {
			return errs.New().SetCode(errs.ErrorNotFound).SetMsg("%s", err)
		}
		return errs.New().SetCode(errs.ErrorInternal).SetMsg("%s", err)
	}
	return nil
}

func (f *Film) IsExist(db DB) *errs.Error {
	var exists bool
	const sqlStr = `SELECT EXISTS(SELECT 1 FROM films WHERE id=$1 LIMIT 1) AS exists`
	err := db.QueryRow(sqlStr, f.ID).Scan(&exists)
	if err != nil {
		return errs.New().SetCode(errs.ErrorInternal).SetMsg("%s", err)
	}
	if !exists {
		return errs.New().SetCode(errs.ErrorNotFound)
	}
	return nil
}

func (f *Film) Convert() *domain.Film {
	releaseDate := f.ReleaseDate.Unix()
	return &domain.Film{
		ID:          f.ID,
		Name:        f.Name,
		ReleaseDate: &releaseDate,
		Duration:    f.Duration,
		Score:       f.Score,
		Comment:     f.Comment,
		Watched:     f.Watched,
	}
}

func LoadFilmByID(db DB, id uint64) (*Film, *errs.Error) {
	f := &Film{}
	const sqlStr = `SELECT * FROM films WHERE id=$1`
	if err := db.QueryRow(sqlStr, id).Scan(&f.ID, &f.UserID, &f.Name, &f.ReleaseDate, &f.Duration, &f.Score, &f.Comment, &f.CreatedAt, &f.UpdatedAt, &f.Watched); err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.New().SetCode(errs.ErrorNotFound).SetMsg("unable to load film: %s", err)
		}
		return nil, errs.New().SetCode(errs.ErrorInternal).SetMsg("unable to load film: %s", err)
	}
	return f, nil
}

func DeleteFilmByID(db DB, userId, id uint64) *errs.Error {
	var exists bool
	errDb := db.QueryRow(`SELECT EXISTS(SELECT 1 FROM films WHERE user_id=$1 AND id=$2) AS exists;`, userId, id).Scan(&exists)
	if errDb != nil && errDb != sql.ErrNoRows {
		return errs.New().SetCode(errs.ErrorInternal).SetMsg("failed to check a film exists: %s", errDb)
	}
	if !exists {
		return errs.New().SetCode(errs.ErrorNotFound).SetMsg("film(id:%d) not found", id)
	}
	const sqlStr = `DELETE FROM films WHERE id=$1;`
	_, err := db.Exec(sqlStr, id)
	if err != nil {
		return errs.New().SetCode(errs.ErrorInternal).SetMsg("failed to delete a film: %s", err)
	}
	return nil
}
