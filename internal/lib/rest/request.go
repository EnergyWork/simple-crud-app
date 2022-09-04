package rest

import (
	"encoding/json"
	"io"
	"net/http"
	errs "simple-crud-app/internal/lib/errors"
)

type Request interface {
	SetHeader(*http.Request)
	SetDigest([]byte)
}

func CreateRequest(r *http.Request, req Request) *errs.Error {
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
	req.SetDigest(reqBts)
	return nil
}
