package models

import (
	"database/sql"
	"time"

	"simple-crud-app/internal/lib/crypto"
	errs "simple-crud-app/internal/lib/errors"
)

type UserAuth struct {
	ID           uint64
	UserName     string
	UserPassword string
	PrivateKey   crypto.PrivateKey // base64 in db
	ExpiredAt    time.Time
}

func (a *UserAuth) Update(db DB) *errs.Error {
	const sqlStr = `UPDATE user_auth SET user_password=$2, private_key=$3, expired_at=$4 WHERE user_name=$1`
	_, errDb := db.Exec(sqlStr, a.UserName, a.UserPassword, a.PrivateKey, a.ExpiredAt)
	if errDb != nil {
		return errs.New().SetCode(errs.ERROR_INTERNAL).SetMsg("failed to update a user: %s", errDb)
	}
	return nil
}

func (a *UserAuth) UserRegister(db DB) *errs.Error {
	// check if user exists
	var exists bool
	if errDb := db.QueryRow(`SELECT EXISTS(SELECT 1 FROM user_auth WHERE user_name = $1) AS exists`, a.UserName).Scan(&exists); errDb != nil {
		return errs.New().SetCode(errs.ERROR_INTERNAL).SetMsg("failed to check a user: %s", errDb)
	}
	if exists {
		return errs.New().SetCode(errs.ERROR_ALREADY_EXISTS).SetMsg("user %s already exists", a.UserName)
	}
	// create user, password already hashed
	sqlStr := `INSERT INTO user_auth(user_name, user_password, private_key, expired_at) VALUES($1, $2, $3, $4)`
	if _, errDb := db.Exec(sqlStr, a.UserName, a.UserPassword, a.PrivateKey, a.ExpiredAt); errDb != nil {
		return errs.New().SetCode(errs.ERROR_INTERNAL).SetMsg("failed to create a new record: %s", errDb)
	}
	return nil
}

func (a *UserAuth) LogIn(db DB) *errs.Error {
	// check if user exists
	var exists bool
	errDb := db.QueryRow(`SELECT EXISTS(SELECT 1 FROM user_auth WHERE user_name = $1) AS exists`, a.UserName).Scan(&exists)
	if errDb != nil && errDb != sql.ErrNoRows {
		return errs.New().SetCode(errs.ERROR_INTERNAL).SetMsg("failed to check a user: %s", errDb)
	}
	if !exists {
		return errs.New().SetCode(errs.ERROR_NOT_FOUND).SetMsg("user %s not found", a.UserName)
	}
	sqlStr := `SELECT private_key, expired_at FROM user_auth WHERE user_name=$1 AND user_password=$2`
	if errDb := db.QueryRow(sqlStr, a.UserName, a.UserPassword).Scan(&a.PrivateKey, &a.ExpiredAt); errDb != nil {
		return errs.New().SetCode(errs.ERROR_INTERNAL).SetMsg("login is failed: %s", errDb)
	}
	return nil
}

func (a *UserAuth) IsExpired() bool {
	return time.Now().After(a.ExpiredAt)
}
