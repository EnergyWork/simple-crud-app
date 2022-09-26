package api

import (
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/lib/rest"
	"simple-crud-app/internal/models"
)

type CustomHeader struct {
	rest.Header
	user *models.User
}

func (obj *CustomHeader) CheckSession(db models.DB) *errs.Error {
	if obj.user.Login != obj.Login { // obj.user.Login loaded by access key  --  obj.Login from X-User header
		return errs.New().SetCode(errs.ERROR_UNAUTHORIZED).SetMsg("unauthorized request")
	}
	session, errApi := obj.user.Session(db)
	if errApi != nil {
		return errs.New().SetCode(errs.ERROR_INTERNAL).SetMsg("%s", errApi)
	}
	if session.IsExpired() {
		return errs.New().SetCode(errs.ERROR_SESSION_EXPIRED).SetMsg("session is expired")
	}
	return nil
}

func (obj *CustomHeader) Authorize(db models.DB) *errs.Error {
	user, errApi := models.LoadUserByLogin(db, obj.Login)
	if errApi != nil {
		return errApi
	}
	obj.user = user // set user to request
	if user.AccessKey != obj.AccessKey {
		return errs.New().SetCode(errs.ERROR_FORBIDDEN).SetMsg("wrong access key")
	}
	return nil
}
