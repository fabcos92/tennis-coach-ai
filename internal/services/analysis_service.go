package services

import (
	"context"
	"encoding/json"
	"strings"
	"tennis-coach-ai/internal/llm"
	"tennis-coach-ai/internal/models"
)

type AnalysisService struct {
	llm llm.Client
}

func NewAnalysisService(llm llm.Client) *AnalysisService {
	return &AnalysisService{llm}
}

func (s *AnalysisService) Analyze(ctx context.Context, req models.AnalyzeRequest) (*models.AnalyzeResponse, error) {
	prompt := buildPrompt(req)

	raw, err := s.llm.Analyze(ctx, prompt)
	if err != nil {
		return nil, err
	}

	clean := extractJSON(raw)
	if clean == "" {
		return fallbackResponse(), nil
	}

	var resp models.AnalyzeResponse
	err = json.Unmarshal([]byte(clean), &resp)
	if err != nil {
		return fallbackResponse(), nil
	}

	if !validate(resp) {
		return fallbackResponse(), nil
	}

	return &resp, nil
}

func extractJSON(raw string) string {
	start := strings.Index(raw, "{")
	end := strings.LastIndex(raw, "}")

	if start == -1 || end == -1 || end <= start {
		return ""
	}

	return raw[start : end+1]
}

func fallbackResponse() *models.AnalyzeResponse {
	return &models.AnalyzeResponse{
		Issues: []string{
			"Unable to analyze data reliably",
		},
		Recommendations: []string{
			"Try again with clearer input",
		},
		FocusArea: "unknown",
	}
}

func validate(resp models.AnalyzeResponse) bool {
	return len(resp.Issues) > 0 && len(resp.Recommendations) > 0
}
