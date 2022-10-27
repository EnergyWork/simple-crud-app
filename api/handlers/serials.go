package handlers

import (
	"database/sql"
	"net/http"
)

type SerialsHandler struct {
	dbConn *sql.DB
}

func NewSerialsHandler(db *sql.DB) *SerialsHandler {
	return &SerialsHandler{dbConn: db}
}

func (obj *SerialsHandler) List(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (obj *SerialsHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (obj *SerialsHandler) Read(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (obj *SerialsHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (obj *SerialsHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
