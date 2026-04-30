package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	models "tennis-coach-ai/internal/models"
	"tennis-coach-ai/internal/services"
)

type AnalyzeHandler struct {
	service *services.AnalysisService
}

func NewAnalyzeHandler(service *services.AnalysisService) *AnalyzeHandler {
	return &AnalyzeHandler{service}
}

func (h *AnalyzeHandler) Analyze(w http.ResponseWriter, r *http.Request) {
	log.Printf("[ANALYZE] request started")

	var req models.AnalyzeRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("[ANALYZE] invalid request body: %v", err)
		writeError(w, http.StatusBadRequest, "INVALID_REQUEST", "invalid request body")
		return
	}

	if req.Type == "" {
		log.Printf("[ANALYZE] validation error: missing type")
		writeError(w, http.StatusBadRequest, "VALIDATION_ERROR", "missing type")
		return
	}

	resp, err := h.service.Analyze(r.Context(), req)
	if err != nil {
		log.Printf("[ANALYZE] service error: %v", err)
		writeError(w, http.StatusInternalServerError, "SERVICE_ERROR", "internal error")
		return
	}

	log.Printf("[ANALYZE] request completed successfully")
	writeJSON(w, http.StatusOK, resp)
}
