package api

import (
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/lib/rest"
	"simple-crud-app/internal/models"
)

type CustomHeader struct {
	rest.Header
}

func (obj *CustomHeader) Authorize(db models.DB) *errs.Error {
	// verify signature
	/*
		load user record by unique userName

	*/

	return nil
}
