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

// ConfigureRouter : all requests are of the POST method
func (s *Server) ConfigureRouter() {
	authHandler := NewAuthHandler(s.GetDB())
	filmsHandler := NewFilmsHandler(s.GetDB())
	/*serialsHandler := NewSerialsHandler(s.GetDB())*/

	router := http.NewServeMux()

	// ping request
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("pong"))
	})

	// authorization handlers
	router.HandleFunc("/auth/registration", authHandler.Registration)
	router.HandleFunc("/auth/login", authHandler.Login)
	router.HandleFunc("/auth/logout", authHandler.Logout)

	// film handlers
	router.HandleFunc("/films/list", filmsHandler.List)
	router.HandleFunc("/films/create", filmsHandler.Create)
	router.HandleFunc("/films/get", filmsHandler.Read)
	router.HandleFunc("/films/update", filmsHandler.Update)
	router.HandleFunc("/films/delete", filmsHandler.Delete)

	// serial handlers
	/*
		someday there will be an implementation
	*/

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
