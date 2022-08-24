package api

import (
	"net/http"
	"simple-crud-app/internal/lib/config"
	"simple-crud-app/internal/lib/rest"
	"simple-crud-app/pkg/database"
)

type Server struct {
	rest.BaseServer
}

func NewHttpServer(cfg *config.Config) (*Server, error) {
	s := &Server{}
	s.InitBase(cfg)
	s.configureRouter()
	return s, nil
}

func (s *Server) configureRouter() {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	s.SetRouter(router)
}

func (s *Server) ConnectToDatabase() error {
	db, err := database.NewPostgresConnection(s.GetConfig())
	s.SetDB(db)
	return err
}
