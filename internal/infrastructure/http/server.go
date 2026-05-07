package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	config "tennis-coach-ai/cfg"
	"tennis-coach-ai/internal/application"
	"tennis-coach-ai/internal/infrastructure/http/handlers"
	"tennis-coach-ai/internal/infrastructure/http/middlewares"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
	HTTP   *http.Server
}

func NewServer(cfg *config.Config, app *application.Application) *Server {
	r := mux.NewRouter()

	s := &Server{
		Router: r,
	}

	r.Use(
		middlewares.CORSMiddleware([]string{"http://localhost:5173"}),
	)

	healthHandler := handlers.NewHealthHandler()
	registerHealthRoutes(r, healthHandler)

	analyzeHandler := handlers.NewAnalyzeHandler(app)
	register(r, analyzeHandler)

	s.HTTP = &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port),
		Handler:      s.Router,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return s
}

func (s *Server) Start() error {
	log.Printf("API listening on %s\n", s.HTTP.Addr)
	return s.HTTP.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("Shutting down API...")
	return s.HTTP.Shutdown(ctx)
}
