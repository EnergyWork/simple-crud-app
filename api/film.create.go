package api

import (
	"net/http"
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/lib/rest"
	"simple-crud-app/internal/models"
	"time"
)

type ReqCreateFilm struct {
	rest.Header
	Name        string     `json:"name"`
	ReleaseDate *time.Time `json:"release_date"`
	Duration    *string    `json:"duration"`
	Score       *uint64    `json:"score"`
	Comment     *string    `json:"comment"`
}

type RplCreateFilm struct {
	rest.Header
}

func (obj *ReqCreateFilm) Validate() *errs.Error {
	if obj.Name == "" {
		return errs.New().SetCode(errs.ERROR_SYNTAX).SetMsg("Name must be not empty")
	}
	return nil
}

// CreateFilm : POST handler for create film requests
func (s *Server) CreateFilm(w http.ResponseWriter, r *http.Request) {
	resp := &RplCreateFilm{}
	// handle request method
	if r.Method != http.MethodPost {
		errApi := errs.New().SetCode(errs.ERROR_METHOD_NOT_ALLOWED).SetMsg("not allowed method - expected POST")
		rest.CreateResponseError(w, resp, errApi)
		return
	}

	req := &ReqCreateFilm{}
	if err := rest.CreateRequest(r, req); err != nil {
		rest.CreateResponseError(w, resp, err)
		return
	}

	// authorization
	if err := req.Authorize(s.GetDB()); err != nil {
		rest.CreateResponseError(w, resp, err)
		return
	}

	film := &models.Film{
		Name:        req.Name,
		ReleaseDate: req.ReleaseDate,
		Duration:    req.Duration,
		Score:       req.Score,
		Comment:     req.Comment,
	}
	if err := film.Create(s.GetDB()); err != nil {
		rest.CreateResponseError(w, resp, err)
		return
	}

	rest.CreateResponse(w, resp)
}
