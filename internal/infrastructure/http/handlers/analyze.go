package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"tennis-coach-ai/internal/application"
	"tennis-coach-ai/internal/application/commands"
	"tennis-coach-ai/internal/domain/input"
	"tennis-coach-ai/internal/infrastructure/http/dto"
	"tennis-coach-ai/internal/infrastructure/http/handlers/shared"
	"tennis-coach-ai/internal/infrastructure/http/mappers"
)

type AnalyzeHandler struct {
	app    *application.Application
	mapper mappers.MatchInputMapper
}

func NewAnalyzeHandler(app *application.Application, mapper mappers.MatchInputMapper) *AnalyzeHandler {
	return &AnalyzeHandler{app, mapper}
}

func (h *AnalyzeHandler) Analyze(w http.ResponseWriter, r *http.Request) {
	log.Printf("[ANALYZE] request started")

	var req dto.AnalyzeRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("[ANALYZE] invalid request body: %v", err)
		shared.WriteError(w, http.StatusBadRequest, "INVALID_REQUEST", "invalid request body")
		return
	}

	inputType, err := h.mapper.ToInputType(req.Type)
	if err != nil {
		log.Printf("[ANALYZE] invalid request body: %v", err)
		shared.WriteError(w, http.StatusBadRequest, "INVALID_REQUEST", "invalid request body")
		return
	}

	var statsInput *input.Stats
	if inputType.IsStats() {
		statsInput, err = h.mapper.ToStats(req.Stats)
		if err != nil {
			log.Printf("[ANALYZE] invalid request body: %v", err)
			shared.WriteError(w, http.StatusBadRequest, "INVALID_REQUEST", "invalid request body")
			return
		}
	}

	var textInput *input.Text
	if inputType.IsText() {
		textInput, err = h.mapper.ToText(req.Text)
		if err != nil {
			log.Printf("[ANALYZE] invalid request body: %v", err)
			shared.WriteError(w, http.StatusBadRequest, "INVALID_REQUEST", "invalid request body")
			return
		}
	}

	command := commands.NewAnalyzeMatchPerformance(inputType, statsInput, textInput)
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
