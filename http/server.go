package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	config "tennis-coach-ai/cfg"
	"tennis-coach-ai/internal/handlers"
	"tennis-coach-ai/internal/llm"
	"tennis-coach-ai/internal/services"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
	HTTP   *http.Server
}

func NewServer(cfg *config.Config) *Server {
	r := mux.NewRouter()

	s := &Server{
		Router: r,
	}

	healhHandler := handlers.NewHealthHandler()
	registerHealthRoutes(r, healhHandler)

	llmClient := llm.NewProvider(llm.Provider(cfg.LLM.Prodiver), cfg.OpenAI.Key)
	analysisService := services.NewAnalysisService(llmClient)
	analyzeHandler := handlers.NewAnalyzeHandler(analysisService)
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
