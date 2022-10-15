package models

import (
	"database/sql"
	"time"

	errs "simple-crud-app/internal/lib/errors"
)

type User struct {
	ID        uint64
	Login     string
	Password  string
	SessionID uint64
	AccessKey string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

func (a *User) Create(db DB) *errs.Error {
	// check if user exists
	var exists bool
	if errDb := db.QueryRow(`SELECT EXISTS(SELECT 1 FROM users WHERE login = $1) AS exists`, a.Login).Scan(&exists); errDb != nil {
		return errs.New().SetCode(errs.ErrorInternal).SetMsg("failed to check a user: %s", errDb)
	}
	if exists {
		return errs.New().SetCode(errs.ErrorAlreadyExists).SetMsg("user with login %s already exists", a.Login)
	}
	// create user, password already hashed
	sqlStr := `INSERT INTO users (login, password, session_id, access_key) VALUES($1, $2, $3, $4)`
	if _, errDb := db.Exec(sqlStr, a.Login, a.Password, a.SessionID, a.AccessKey); errDb != nil {
		return errs.New().SetCode(errs.ErrorInternal).SetMsg("failed to create a new record: %s", errDb)
	}
	return nil
}

func (a *User) LogIn(db DB) *errs.Error {
	// check if user exists
	var exists bool
	errDb := db.QueryRow(`SELECT EXISTS(SELECT 1 FROM user_auth WHERE user_name = $1) AS exists`, a.Login).Scan(&exists)
	if errDb != nil && errDb != sql.ErrNoRows {
		return errs.New().SetCode(errs.ErrorInternal).SetMsg("failed to check a user: %s", errDb)
	}
	if !exists {
		return errs.New().SetCode(errs.ErrorNotFound).SetMsg("user %s not found", a.Login)
	}
	sqlStr := `SELECT session_id, access_key FROM users WHERE login=$1 AND password=$2`
	if errDb := db.QueryRow(sqlStr, a.Login, a.Password).Scan(a.ID, a.SessionID, a.AccessKey); errDb != nil {
		return errs.New().SetCode(errs.ErrorInternal).SetMsg("login is failed: %s", errDb)
	}
	return nil
}

func (a *User) Session(db DB) (*Session, *errs.Error) {
	s, err := LoadSession(db, a.SessionID)
	if err != nil {
		return nil, errs.New().SetCode(errs.ErrorInternal).SetMsg("%s", err)
	}
	return s, nil
}

func (a *User) CloseSession(db DB) *errs.Error {
	errApi := CloseSession(db, a.SessionID)
	if errApi != nil {
		return errApi
	}
	return nil
}

func (a *User) ChangePassword(db DB) *errs.Error {
	//! TODO implement me
	return nil
}

func LoadUserByLogin(db DB, login string) (*User, *errs.Error) {
	u := &User{}
	const sqlStr = `SELECT * FROM users WHERE login=$1`
	err := db.QueryRow(sqlStr, login).Scan(&u.ID, &u.SessionID, &u.Login, &u.Password, &u.AccessKey, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.New().SetCode(errs.ErrorNotFound).SetMsg("%s", err)
		}
		return nil, errs.New().SetCode(errs.ErrorInternal).SetMsg("%s", err)
	}
	return u, nil
}

func LoadUserByAccessKey(db DB, accessKey string) (*User, *errs.Error) {
	u := &User{}
	const sqlStr = `SELECT * FROM users WHERE access_key=$1`
	err := db.QueryRow(sqlStr, accessKey).Scan(&u.ID, &u.SessionID, &u.Login, &u.Password, &u.AccessKey, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, errs.New().SetCode(errs.ErrorInternal).SetMsg("%s", err)
	}
	return u, nil
}
