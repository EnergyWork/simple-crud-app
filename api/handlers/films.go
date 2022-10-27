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

func (obj *FilmsHandler) Read(w http.ResponseWriter, r *http.Request) {
	_ = rest.BaseHandler(w, r, &films.ReqFilmGet{}, &films.RespFilmGet{}, obj.dbConn, "FilmGet", true)
}

func (obj *FilmsHandler) Update(w http.ResponseWriter, r *http.Request) {
	_ = rest.BaseHandler(w, r, &films.ReqFilmUpdate{}, &films.RespFilmUpdate{}, obj.dbConn, "FilmUpdate", true)
}

func (obj *FilmsHandler) Delete(w http.ResponseWriter, r *http.Request) {
	_ = rest.BaseHandler(w, r, &films.ReqFilmDelete{}, &films.RespFilmDelete{}, obj.dbConn, "FilmDelete", true)
}
