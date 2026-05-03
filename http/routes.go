package http

import (
	"net/http"
	"tennis-coach-ai/internal/handlers"

	"github.com/gorilla/mux"
)

func registerHealthRoutes(r *mux.Router, healthHandler *handlers.HealthHandler) {
	r.HandleFunc("/healthz", healthHandler.Check).Methods(http.MethodGet, http.MethodOptions)
}

func register(r *mux.Router, userHandler *handlers.AnalyzeHandler) {
	r.HandleFunc("/analyze", userHandler.Analyze).Methods(http.MethodPost, http.MethodOptions)
}
