package api

import (
	"net/http"
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/lib/logger"
	"simple-crud-app/internal/lib/rest"
	"simple-crud-app/internal/models"
	"time"
)

type ReqCreateFilm struct {
	CustomHeader
	Name        string     `json:"name"`
	ReleaseDate *time.Time `json:"release_date"`
	Duration    *string    `json:"duration"`
	Score       *uint64    `json:"score"`
	Comment     *string    `json:"comment"`
}

type RespCreateFilm struct {
	rest.Header
}

func (obj *ReqCreateFilm) Validate() *errs.Error {
	if obj.Name == "" {
		return errs.New().SetCode(errs.ErrorRequestSyntax).SetMsg("Name must be not empty")
	}
	return nil
}

// CreateFilm : POST handler for create film requests
func (s *Server) CreateFilm(w http.ResponseWriter, r *http.Request) {
	l := logger.NewLogger().SetMethod("CreateFilm")
	req := &ReqCreateFilm{}
	resp := &RespCreateFilm{}

	//unmarshal input request into struct
	if errApi := rest.CreateRequest(r, req, http.MethodPost); errApi != nil {
		rest.CreateResponseError(w, resp, errApi)
		l.Errorf("errro: unable create request - %s", errApi)
		return
	}

	if errApi := rest.Prepare(s.GetDB(), req); errApi != nil {
		rest.CreateResponseError(w, resp, errApi)
		l.Errorf("error: %s", errApi)
		return
	}

	l.Debugf("->REQ: %+v", req)

	//? ///////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// business logic
	film := &models.Film{
		Name:        req.Name,
		UserID:      req.user.ID,
		ReleaseDate: req.ReleaseDate,
		Duration:    req.Duration,
		Score:       req.Score,
		Comment:     req.Comment,
	}
	if err := film.Create(s.GetDB()); err != nil {
		rest.CreateResponseError(w, resp, err)
		l.Errorf("error: %s", err)
		return
	}

	//? ///////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// response
	rest.CreateResponse(w, resp)
	l.Debugf("<-RESP: %+v", resp)
}
