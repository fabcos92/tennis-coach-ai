package handlers

import (
	"encoding/json"
	"net/http"
)

type AnalyzeHandler struct{}

func NewAnalyzeHandler() *AnalyzeHandler {
	return &AnalyzeHandler{}
}

func (h *AnalyzeHandler) Analyze(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{
		"status": "ok",
	})
}
