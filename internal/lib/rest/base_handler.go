package rest

import (
	"database/sql"
	"github.com/google/uuid"
	"net/http"
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/lib/logger"
)

func BaseHandler(w http.ResponseWriter, r *http.Request, req Request, resp Response, dbConn *sql.DB, method string, auth bool) *errs.Error {
	req.SetReqID(uuid.NewString())
	l := logger.NewLogger().SetMethod(method)

	var errApi *errs.Error

	// unmarshal input request into struct and setup request data
	if errApi = CreateRequest(r, dbConn, req, auth); errApi != nil {
		l.Errorf("unable to create a request - %s", errApi)
		CreateResponseError(w, errApi)
		return errApi
	}

	// then Execute business logic
	if resp, errApi = req.Execute(); errApi != nil {
		l.Errorf("execute returns: %s", errApi)
		CreateResponseError(w, errApi)
		return errApi
	}

	CreateResponse(w, resp)
	return nil
}
