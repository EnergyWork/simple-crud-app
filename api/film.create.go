package api

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/lib/rest"
	"simple-crud-app/internal/models"
	"time"
)

type ReqCreateFilm struct {
	Type        string
	Name        string
	ReleaseDate time.Time
	Duration    string
	SerialCount int64
	Score       uint64
	Comment     string
}

func (obj *ReqCreateFilm) Validate() *errs.Error {
	return nil
}

// CreateFilm : POST handler for create film requests
func (s *Server) CreateFilm(w http.ResponseWriter, r *http.Request) {
	// handle request method
	if r.Method != http.MethodPost {
		errApi := errs.New().SetCode(errs.ERROR_METHOD_NOT_ALLOWED).SetMsg("not allowed method - expected POST")
		rest.CreateRplError(w, errApi)
		return
	}
	// unmarshal request body
	reqBts, err := io.ReadAll(r.Body)
	if err != nil {
		errApi := errs.New().SetCode(errs.ERROR_INTERNAL)
		rest.CreateRplError(w, errApi)
		return
	}

	req := &ReqCreateFilm{}
	if err := json.Unmarshal(reqBts, &req); err != nil {
		errApi := errs.New().SetCode(errs.ERROR_INTERNAL)
		rest.CreateRplError(w, errApi)
		return
	}

	film := &models.Film{
		Name:        req.Name,
		Type:        sql.NullString{String: req.Type},
		ReleaseDate: sql.NullTime{Time: req.ReleaseDate},
		Duration:    sql.NullString{String: req.Duration},
		SerialCount: sql.NullInt64{Int64: req.SerialCount},
		Score:       req.Score,
		Comment:     sql.NullString{String: req.Comment},
	}
	if err := film.Create(s.GetDB()); err != nil {
		rest.CreateRplError(w, err)
		return
	}
}
