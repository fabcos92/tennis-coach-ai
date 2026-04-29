package services

import (
	"context"
	"encoding/json"
	"tennis-coach-ai/internal/llm"
	"tennis-coach-ai/internal/models"
)

type AnalysisService struct {
	llm *llm.Client
}

func (s *AnalysisService) Analyze(ctx context.Context, req models.AnalyzeRequest) (*models.AnalyzeResponse, error) {
	prompt := BuildPrompt(req)

	raw, err := s.llm.Analyze(prompt)
	if err != nil {
		return nil, err
	}

	var resp models.AnalyzeResponse
	err = json.Unmarshal([]byte(raw), &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
