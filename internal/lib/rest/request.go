package rest

import (
	"encoding/json"
	"io"
	"net/http"
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/models"
)

type Request interface {
	SetHeader(*http.Request)
	SetError(err *errs.Error)
	Authorize(db models.DB) *errs.Error
	CheckSession(db models.DB) *errs.Error
	Validate() *errs.Error
}

func CreateRequest(r *http.Request, req Request, expectedMethod string) *errs.Error {
	// handle request method
	if r.Method != http.MethodPost {
		return errs.New().SetCode(errs.ERROR_METHOD_NOT_ALLOWED).SetMsg("not allowed method - expected POST")
	}
	// read request body
	reqBts, err := io.ReadAll(r.Body)
	if err != nil {
		return errs.New().SetCode(errs.ERROR_INTERNAL).SetMsg("internal system error")
	}
	// unmarshal bytes to request struct
	if err := json.Unmarshal(reqBts, &req); err != nil {
		return errs.New().SetCode(errs.ERROR_INTERNAL).SetMsg("internal system error")
	}
	req.SetHeader(r)
	return nil
}
