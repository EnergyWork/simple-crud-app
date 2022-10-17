package handlers

import (
	"net/http"
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/lib/logger"
	"simple-crud-app/internal/lib/rest"
	"simple-crud-app/internal/models"
)

type ReqFilmList struct {
	CustomHeader
	// TODO SortParameters
	// TODO VALIDATE
	Offset uint64
	Limit  uint64
}

type RespFilmList struct {
	rest.Header
	Films []models.Film
	Total uint64
}

func (obj *ReqFilmList) Validate() *errs.Error {
	return nil
}

// FilmList :
func (s *Server) FilmList(w http.ResponseWriter, r *http.Request) {
	l := logger.NewLogger().SetMethod("FilmList")
	req := &ReqFilmList{}
	resp := &RespFilmList{}

	//unmarshal input request into struct
	if err := rest.CreateRequest(r, &s.BaseServer, req, http.MethodPost); err != nil {
		rest.CreateResponseError(w, err)
		l.Errorf("error: unable create request - %s", err)
		return
	}

	if errApi := rest.Prepare(s.GetDB(), req); errApi != nil {
		rest.CreateResponseError(w, errApi)
		l.Errorf("error: %s", errApi)
		return
	}

	l.Debugf("->REQ: %+v", req)

	//? business logic //////////////////////////////////////////////////////////////////////////////////////////////////////////

	filmList := models.FilmList{
		UserID: req.user.ID,
		Offset: req.Offset,
		Limit:  req.Limit,
	}
	list, total, errApi := filmList.GetList(s.GetDB())
	if errApi != nil {
		rest.CreateResponseError(w, errApi)
		l.Error(errApi)
		return
	}

	//? response /////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	resp.Films = list
	resp.Total = total
	rest.CreateResponse(w, resp)
	l.Debugf("<-RESP: %+v", resp)
}
