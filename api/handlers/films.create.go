package handlers

import (
	"github.com/google/uuid"
	"net/http"
	"simple-crud-app/api/usecase/films"
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/lib/logger"
	"simple-crud-app/internal/lib/rest"
)

// CreateFilm : POST handler for create film requests
func (s *Server) CreateFilm(w http.ResponseWriter, r *http.Request) {
	req := &films.ReqCreateFilm{}

	req.SetReqID(uuid.NewString())
	l := logger.NewLogger().SetMethod("CreateFilm")

	var errApi *errs.Error

	//unmarshal input request into struct
	if errApi = rest.CreateRequest(r, s.GetDB(), req, true); errApi != nil {
		rest.CreateResponseError(w, errApi)
		l.Errorf("error: unable create request - %s", errApi)
		return
	}

	// then Execute business logic
	resp, err := req.Execute()
	if errApi != nil {
		rest.CreateResponseError(w, err)
		return
	}

	rest.CreateResponse(w, resp)
}
