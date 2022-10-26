package handlers

import (
	"database/sql"
	"net/http"
	"simple-crud-app/api/usecase/films"
	"simple-crud-app/internal/lib/rest"
)

type FilmsHandler struct {
	dbConn *sql.DB
}

func NewFilmsHandler(db *sql.DB) *FilmsHandler {
	return &FilmsHandler{dbConn: db}
}

func (obj *FilmsHandler) List(w http.ResponseWriter, r *http.Request) {
	_ = rest.BaseHandler(w, r, &films.ReqFilmList{}, &films.RespFilmList{}, obj.dbConn, "FilmsList", true)
}

func (obj *FilmsHandler) Create(w http.ResponseWriter, r *http.Request) {
	_ = rest.BaseHandler(w, r, &films.ReqCreateFilm{}, &films.RespCreateFilm{}, obj.dbConn, "FilmCreate", true)
}
