package api

import (
	"net/http"
	"simple-crud-app/internal/lib/config"
	"simple-crud-app/internal/lib/rest"
	"simple-crud-app/pkg/database"

	_ "github.com/lib/pq"
)

type Server struct {
	rest.BaseServer
}

func NewHttpServer(cfg *config.Config) *Server {
	s := &Server{}
	s.InitBase(cfg)
	s.configureRouter()
	return s
}

func (s *Server) configureRouter() {
	router := http.NewServeMux()
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	})
	router.HandleFunc("/auth/registration", s.AuthRegister) // [POST]
	router.HandleFunc("/auth/login", s.AuthLogin) // [POST]
	s.SetRouter(router)
}

func (s *Server) ConnectToDatabase() error {
	db, err := database.NewPostgresConnection(s.GetConfig())
	if err != nil {
		return err
	}
	s.SetDB(db)
	return nil
}
