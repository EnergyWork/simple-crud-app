package usecase

import (
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/lib/rest"
	"simple-crud-app/internal/models"
)

type CustomHeader struct {
	rest.Header
	User *models.User
}

func (obj *CustomHeader) CheckSession(db models.DB) *errs.Error {
	if obj.User.Login != obj.Login { // obj.User.Login loaded by access key  --  obj.Login from X-User header
		return errs.New().SetCode(errs.ErrorUnauthorized).SetMsg("unauthorized request")
	}
	session, errApi := obj.User.Session(db)
	if errApi != nil {
		return errs.New().SetCode(errs.ErrorInternal).SetMsg("%s", errApi)
	}
	if session.IsExpired() {
		return errs.New().SetCode(errs.ErrorSessionExpired).SetMsg("session is expired")
	}
	return nil
}

func (obj *CustomHeader) Authorize(db models.DB) *errs.Error {
	user, errApi := models.LoadUserByLogin(db, obj.Login)
	if errApi != nil {
		return errApi
	}
	obj.User = user // setup User for request
	if user.AccessKey != obj.AccessKey {
		return errs.New().SetCode(errs.ErrorForbidden).SetMsg("wrong access key")
	}
	return nil
}
