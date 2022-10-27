package handlers

import (
	"database/sql"
	"net/http"

	"simple-crud-app/api/usecase/auth"
	"simple-crud-app/internal/lib/rest"
)

type AuthHandler struct {
	dbConn *sql.DB
}

func NewAuthHandler(db *sql.DB) *AuthHandler {
	return &AuthHandler{dbConn: db}
}

// Registration : регистрация нового пользователя в системе
func (h AuthHandler) Registration(w http.ResponseWriter, r *http.Request) {
	_ = rest.BaseHandler(w, r, &auth.ReqAuthRegister{}, &auth.RespAuthRegister{}, h.dbConn, "Registration", false)
}

// Login : авторизация пользователя в системе
func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	_ = rest.BaseHandler(w, r, &auth.ReqAuthLogin{}, &auth.RespAuthLogin{}, h.dbConn, "Login", false)
}

// Logout : выход пользователя из системы
func (h AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	_ = rest.BaseHandler(w, r, &auth.ReqAuthLogout{}, &auth.RespAuthLogout{}, h.dbConn, "Logout", true)
}
