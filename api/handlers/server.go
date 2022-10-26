package handlers

import (
	"net/http"
	"simple-crud-app/internal/lib/config"
	"simple-crud-app/internal/lib/rest"
	"simple-crud-app/internal/lib/rest/middleware"
	"simple-crud-app/pkg/database"

	_ "github.com/lib/pq"
)

var s *Server

type Server struct {
	rest.BaseServer
}

func NewHttpServer(cfg *config.Config) *Server {
	s = &Server{}
	s.InitBase(cfg)
	return s
}

func (s *Server) ConfigureRouter() {
	authHandler := NewAuthHandler(s.GetDB())
	filmsHandler := NewFilmsHandler(s.GetDB())

	router := http.NewServeMux()

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("pong"))
	})

	router.HandleFunc("/auth/registration", authHandler.Registration) // [POST]
	router.HandleFunc("/auth/login", authHandler.Login)               // [POST]
	router.HandleFunc("/auth/logout", authHandler.Logout)             // [POST]

	router.HandleFunc("/films/list", filmsHandler.List)     // [POST]
	router.HandleFunc("/films/create", filmsHandler.Create) // [POST]
	router.HandleFunc("/films/delete", s.FilmDelete)        // [POST]
	router.HandleFunc("/films/update", s.FilmUpdate)        // [POST]

	wrappedMux := middleware.NewLoggerRequest(router)

	s.SetRouter(wrappedMux)
}

func (s *Server) ConnectToDatabase() error {
	db, err := database.NewPostgresConnection(s.GetConfig())
	if err != nil {
		return err
	}
	s.SetDB(db)
	return nil
}
