package rest

import (
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/models"
)

func Prepare(db models.DB, req Request, auth bool) *errs.Error {
	// if authorized request
	if auth {
		// authorization
		if errApi := req.Authorize(db); errApi != nil {
			return errApi
		}
		// check session
		if errApi := req.CheckSession(db); errApi != nil {
			return errApi
		}
	}
	// validation
	if errApi := req.Validate(); errApi != nil {
		return errApi
	}
	return nil
}
