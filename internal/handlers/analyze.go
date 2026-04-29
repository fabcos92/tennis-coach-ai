package handlers

import (
	"encoding/json"
	"net/http"
	models "tennis-coach-ai/internal/models"
	"tennis-coach-ai/internal/services"
)

type AnalyzeHandler struct {
	service services.AnalysisService
}

func NewAnalyzeHandler() *AnalyzeHandler {
	return &AnalyzeHandler{}
}

func (h *AnalyzeHandler) Analyze(w http.ResponseWriter, r *http.Request) {
	var req models.AnalyzeRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if req.Type == "" {
		http.Error(w, "missing type", http.StatusBadRequest)
		return
	}

	resp, err := h.service.Analyze(r.Context(), req)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, resp)
}
