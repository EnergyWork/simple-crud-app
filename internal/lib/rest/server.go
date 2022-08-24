package rest

import (
	"database/sql"
	"fmt"
	"net/http"
	"simple-crud-app/internal/lib/config"
)

type BaseServer struct {
	server *http.Server
	router *http.ServeMux
	config *config.Config
	db     *sql.DB
}

func (bs *BaseServer) InitBase(cfg *config.Config) {
	bs.config = cfg
	bs.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", bs.config.Api.Host, bs.config.Api.Port),
		Handler: bs.router,
	}
}

func (bs *BaseServer) Run() error {
	return bs.server.ListenAndServe()
}

func (bs *BaseServer) SetDB(db *sql.DB) {
	bs.db = db
}

func (bs *BaseServer) SetRouter(router *http.ServeMux) {
	bs.router = router
}

func (bs *BaseServer) GetConfig() *config.Config {
	return bs.config
}
