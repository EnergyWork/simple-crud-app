package rest

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"net/http"
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/models"
)

type Request interface {
	SetHeader(*http.Request, *sql.DB)
	SetError(*errs.Error)
	SetReqID(string)
	GetReqID() string
	Authorize(models.DB) *errs.Error
	CheckSession(models.DB) *errs.Error
	Validate() *errs.Error
	Execute() (Response, *errs.Error)
}

func CreateRequest(r *http.Request, dbConn *sql.DB, req Request, authorize bool) *errs.Error {
	// read request body
	reqBts, err := io.ReadAll(r.Body)
	if err != nil {
		return errs.New().SetCode(errs.ErrorInternal).SetMsg("internal system error: %s", err)
	}

	// unmarshal bytes to request struct
	if err := json.Unmarshal(reqBts, &req); err != nil {
		return errs.New().SetCode(errs.ErrorRequestSyntax).SetMsg("unmarshal body to request struct: %s", err)
	}

	req.SetHeader(r, dbConn)
	req.SetReqID(uuid.NewString())

	// auth, session, validate
	if errApi := Prepare(dbConn, req, authorize); errApi != nil {
		return errApi
	}

	return nil
}
