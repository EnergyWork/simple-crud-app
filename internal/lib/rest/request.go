package rest

import (
	"encoding/json"
	"io"
	"net/http"
	errs "simple-crud-app/internal/lib/errors"
)

type Request interface {
	SetKey(string)
}

const (
	XAuth = "X-Auth"
)

func CreateRequest(r *http.Request, req Request) *errs.Error {
	// getting authorize key
	key := r.Header.Get(XAuth)
	if key == "" {
		return errs.New().SetCode(errs.ERROR_UNAUTHORIZED)
	}
	req.SetKey(key)
	// read request body
	reqBts, err := io.ReadAll(r.Body)
	if err != nil {
		return errs.New().SetCode(errs.ERROR_INTERNAL).SetMsg("internal system error")
	}
	// unmarshal bytes to request struct
	if err := json.Unmarshal(reqBts, &req); err != nil {
		return errs.New().SetCode(errs.ERROR_INTERNAL).SetMsg("internal system error")
	}
	return nil
}
