package rest

import (
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"net/http"
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/models"
)

type Request interface {
	SetHeader(*http.Request)
	SetError(*errs.Error)
	SetReqID(string)
	Authorize(models.DB) *errs.Error
	CheckSession(models.DB) *errs.Error
	Validate() *errs.Error
}

func CreateRequest(r *http.Request, req Request, expectedMethod string) *errs.Error {
	// handle request method
	if r.Method != expectedMethod {
		return errs.New().SetCode(errs.ErrorMethodNotAllowed).SetMsg("not allowed method - expected POST")
	}
	// read request body
	reqBts, err := io.ReadAll(r.Body)
	if err != nil {
		return errs.New().SetCode(errs.ErrorInternal).SetMsg("internal system error: read request body")
	}
	// unmarshal bytes to request struct
	if err := json.Unmarshal(reqBts, &req); err != nil {
		return errs.New().SetCode(errs.ErrorInternal).SetMsg("internal system error: unmarshal body to request struct")
	}
	req.SetHeader(r)
	req.SetReqID(uuid.NewString())
	return nil
}
