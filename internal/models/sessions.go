package models

import (
	errs "simple-crud-app/internal/lib/errors"
	"time"

	"github.com/google/uuid"
)

const (
	DefaultTTL = 24 * time.Hour
)

type Session struct {
	ID       uint64
	Token    string
	Created  time.Time
	Deadline time.Time
}

func (s *Session) UpdateTTL(db DB, ttl ...time.Duration) *errs.Error {
	newToken := uuid.NewString()
	if ttl != nil {
		s.Deadline = time.Now().Add(ttl[0])
	} else {
		s.Deadline = time.Now().Add(DefaultTTL)
	}
	s.Token = newToken
	const sqlStr = `UPDATE sessions SET token=$1, deadline=$2 WHERE id=$3`
	_, errApi := db.Exec(sqlStr, s.Token, s.Deadline, s.ID)
	if errApi != nil {
		return errs.New().SetCode(errs.ErrorInternal)
	}
	return nil
}

func (s *Session) IsExpired() bool {
	return time.Now().After(s.Deadline)
}

func NewSession(db DB, ttl ...time.Duration) (*Session, *errs.Error) {
	s := &Session{}
	s.Token = uuid.NewString()
	s.Created = time.Now()
	var deadline time.Time
	if ttl != nil {
		deadline = time.Now().Add(ttl[0])
	} else {
		deadline = time.Now().Add(DefaultTTL)
	}
	s.Deadline = deadline
	const sqlStr = `INSERT INTO sessions (token, created, deadline) VALUES($1,$2,$3) RETURNING id`
	err := db.QueryRow(sqlStr, s.Token, s.Created, s.Deadline).Scan(&s.ID)
	if err != nil {
		return nil, errs.New().SetCode(errs.ErrorInternal).SetMsg(err.Error())
	}
	return s, nil
}

func LoadSession(db DB, id uint64) (*Session, *errs.Error) {
	s := &Session{}
	const sqlStr = `SELECT * FROM sessions WHERE id=$1`
	err := db.QueryRow(sqlStr, id).Scan(&s.ID, &s.Token, &s.Created, &s.Deadline)
	if err != nil {
		return nil, errs.New().SetCode(errs.ErrorInternal).SetMsg("%s", err)
	}
	return s, nil
}

func CloseSession(db DB, id uint64) *errs.Error {
	const sqlStr = `UPDATE sessions SET deadline=$2 WHERE id=$1`
	_, err := db.Exec(sqlStr, id, time.Now())
	if err != nil {
		return errs.New().SetCode(errs.ErrorInternal).SetMsg("%s", err)
	}
	return nil
}
