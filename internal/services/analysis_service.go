package services

import (
	"context"
	"encoding/json"
	"fmt"
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

	resp, err := parseAndValidateLLMResponse(raw)
	if err != nil {
		return fallbackResponse(), nil
	}

	if err := validateAgainstInput(req, resp); err != nil {
		return fallbackResponse(), nil
	}

	return resp, nil
}

func parseAndValidateLLMResponse(raw string) (*models.AnalyzeResponse, error) {
	var resp *models.AnalyzeResponse

	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		return &models.AnalyzeResponse{}, fmt.Errorf("invalid JSON: %w", err)
	}

	if resp.FocusArea == "" {
		return &models.AnalyzeResponse{}, fmt.Errorf("missing focus_area")
	}

	if resp.Issues == nil {
		resp.Issues = []models.Issue{}
	}

	if resp.Recommendations == nil {
		resp.Recommendations = []string{}
	}

	return resp, nil
}

func validateAgainstInput(req models.AnalyzeRequest, resp *models.AnalyzeResponse) error {
	for _, issue := range resp.Issues {
		lower := strings.ToLower(issue.Text)

		if strings.Contains(lower, "unforced error") {
			if req.Stats.UnforcedErrors <= 3 && strings.Contains(lower, "high") {
				return fmt.Errorf("invalid issue: exaggeration detected")
			}
		}
	}

	return nil
}

func fallbackResponse() *models.AnalyzeResponse {
	return &models.AnalyzeResponse{
		Issues: []models.Issue{{
			Text:     "Unable to analyze data reliably",
			Severity: "low",
		}},
		Recommendations: []string{
			"Try again with clearer input",
		},
		FocusArea: "unknown",
	}
}
