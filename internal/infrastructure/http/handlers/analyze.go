package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"tennis-coach-ai/internal/application"
	"tennis-coach-ai/internal/application/commands"
	"tennis-coach-ai/internal/infrastructure/http/dto"
	"tennis-coach-ai/internal/infrastructure/http/handlers/shared"
)

type AnalyzeHandler struct {
	app *application.Application
}

func NewAnalyzeHandler(app *application.Application) *AnalyzeHandler {
	return &AnalyzeHandler{app}
}

func (h *AnalyzeHandler) Analyze(w http.ResponseWriter, r *http.Request) {
	log.Printf("[ANALYZE] request started")

	var req dto.AnalyzeRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("[ANALYZE] invalid request body: %v", err)
		shared.WriteError(w, http.StatusBadRequest, "INVALID_REQUEST", "invalid request body")
		return
	}

	if req.Type == "" {
		log.Printf("[ANALYZE] validation error: missing type")
		shared.WriteError(w, http.StatusBadRequest, "VALIDATION_ERROR", "missing type")
		return
	}

	command := commands.NewAnalyzeMatchPerformance(req.Type, req.Text, req.Stats.FirstServeInPct, req.Stats.SecondServeInPct, req.Stats.UnforcedErrors)
	resp, err := h.app.Commands.AnalyzeMatchPerformance.Execute(r.Context(), command)
	if err != nil {
		log.Printf("[ANALYZE] service error: %v", err)
		shared.WriteError(w, http.StatusInternalServerError, "SERVICE_ERROR", "internal error")
		return
	}

	var issues = []dto.Issue{}
	for _, issue := range resp.Issues {
		issues = append(issues, dto.NewIssue(issue.Text, issue.Severity.String()))
	}
	output := dto.NewAnalyzeResponse(issues, resp.Recommendations, resp.FocusArea)

	log.Printf("[ANALYZE] request completed successfully")
	shared.WriteJSON(w, http.StatusOK, output)
}
