package rest

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"log"
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
		return errs.New().SetCode(errs.ErrorInternal).SetMsg("internal system error: read request body")
	}
	// unmarshal bytes to request struct
	if err := json.Unmarshal(reqBts, &req); err != nil {
		log.Println(string(reqBts))
		return errs.New().SetCode(errs.ErrorInternal).SetMsg("internal system error: unmarshal body to request struct")
	}

	req.SetHeader(r, dbConn)
	req.SetReqID(uuid.NewString())

	// auth, session, validate
	if errApi := Prepare(dbConn, req, authorize); errApi != nil {
		return errApi
	}

	return nil
}
