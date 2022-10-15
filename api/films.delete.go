package api

import (
	"net/http"
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/lib/logger"
	"simple-crud-app/internal/lib/rest"
	"simple-crud-app/internal/models"
)

type ReqFilmDelete struct {
	CustomHeader
	ID uint64 `json:"film_id"`
}

type RespFilmDelete struct {
	rest.Header
}

func (obj *ReqFilmDelete) Validate() *errs.Error {
	if obj.ID != 0 {
		return errs.New().SetCode(errs.ErrorRequestSyntax).SetMsg("ID must be not null")
	}
	return nil
}

// FilmDelete :
func (s *Server) FilmDelete(w http.ResponseWriter, r *http.Request) {
	l := logger.NewLogger().SetMethod("FilmDelete")
	req := &ReqFilmDelete{}
	resp := &RespFilmDelete{}

	//unmarshal input request into struct
	if err := rest.CreateRequest(r, req, http.MethodPost); err != nil {
		rest.CreateResponseError(w, resp, err)
		l.Errorf("errro: unable create request - %s", err)
		return
	}

	if errApi := rest.Prepare(s.GetDB(), req); errApi != nil {
		rest.CreateResponseError(w, resp, errApi)
		l.Errorf("error: %s", errApi)
		return
	}

	l.Debugf("->REQ: %+v", req)

	//? ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// business logic

	if errApi := models.DeleteFilmByID(s.GetDB(), req.user.ID, req.ID); errApi != nil {
		rest.CreateResponseError(w, resp, errApi)
		l.Error(errApi)
		return
	}

	//? ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// response
	rest.CreateResponse(w, resp)
	l.Debugf("<-RESP: %+v", resp)
}
