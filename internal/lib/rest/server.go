package rest

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"simple-crud-app/internal/lib/config"
	"time"
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
		Addr: fmt.Sprintf("%s:%s", bs.config.Api.Host, bs.config.Api.Port),
	}
}

func (bs *BaseServer) Run() {
	go func() {
		if err := bs.server.ListenAndServe(); err != nil {
			return
		}
	}()
	time.Sleep(10 * time.Millisecond)
}

func (bs *BaseServer) SetDB(db *sql.DB) {
	bs.db = db
}
func (bs *BaseServer) GetDB() *sql.DB {
	return bs.db
}

func (bs *BaseServer) SetRouter(router *http.ServeMux) {
	bs.router = router
	bs.server.Handler = router
}

func (bs *BaseServer) GetConfig() *config.Config {
	return bs.config
}

func (bs *BaseServer) Shutdown(ctx context.Context) error {
	return bs.server.Shutdown(ctx)
}
