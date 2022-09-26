package api

import (
	"net/http"
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/lib/logger"
	"simple-crud-app/internal/lib/rest"
	"simple-crud-app/internal/models"
)

type ReqFilmUpdate struct {
	CustomHeader
	Film models.Film `json:"updated_film"`
}

type RespFilmUpdate struct {
	rest.Header
}

func (obj *ReqFilmUpdate) Validate() *errs.Error {
	if obj.Film.ID == 0 {
		return errs.New().SetCode(errs.ERROR_SYNTAX).SetMsg("Film.ID must be not null")
	}
	return nil
}

// FilmUpdate :
func (s *Server) FilmUpdate(w http.ResponseWriter, r *http.Request) {
	l := logger.NewLogger().SetMethod("FilmUpdate")
	req := &ReqFilmUpdate{}
	resp := &RespFilmUpdate{}

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

	if errApi := req.Film.Update(s.GetDB()); errApi != nil {
		rest.CreateResponseError(w, resp, errApi)
		l.Error(errApi)
		return
	}

	//? ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// response
	rest.CreateResponse(w, resp)
	l.Debugf("<-RESP: %+v", resp)
}
