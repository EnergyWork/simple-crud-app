package rest

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"simple-crud-app/internal/lib/config"
	errs "simple-crud-app/internal/lib/errors"
	"time"
)

type BaseServer struct {
	server  *http.Server
	router  *http.ServeMux
	config  *config.Config
	db      *sql.DB
	errDesc *errs.ErrorDescription
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

func (bs *BaseServer) SetRouter(router http.Handler) {
	bs.server.Handler = router
}

func (bs *BaseServer) GetConfig() *config.Config {
	return bs.config
}

func (bs *BaseServer) Shutdown(ctx context.Context) error {
	return bs.server.Shutdown(ctx)
}

func (bs *BaseServer) InitErrors(filePath string) error {
	desc := errs.NewErrorDescription(filePath)
	if desc == nil {
		return errors.New("unable to init errors descriptions")
	}
	return nil
}

func (bs *BaseServer) Errors() *errs.ErrorDescription {
	return bs.errDesc
}
