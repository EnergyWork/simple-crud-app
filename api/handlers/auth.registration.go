package handlers

import (
	"github.com/google/uuid"
	"net/http"
	"simple-crud-app/api/requests/auth"
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/lib/logger"
	"simple-crud-app/internal/lib/rest"
)

// Registration : POST handler for user registration
func Registration(w http.ResponseWriter, r *http.Request) {
	req := &auth.ReqAuthRegister{}
	resp := &auth.RespAuthRegister{}

	req.SetReqID(uuid.NewString())
	l := logger.NewLogger().SetMethod("Registration").SetID(req.ReqID)

	var errApi *errs.Error

	// unmarshal input request into struct and setup request data
	if errApi = rest.CreateRequest(r, &s.BaseServer, req, http.MethodPost); errApi != nil {
		rest.CreateResponseError(w, errApi)
		l.Errorf("error: unable create request - %s", errApi)
		return
	}

	if resp, errApi = req.Execute(); errApi != nil {
		rest.CreateResponseError(w, errApi)
		return
	}

	rest.CreateResponse(w, resp)
}
