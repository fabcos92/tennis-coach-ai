package services

import (
	"context"
	"tennis-coach-ai/internal/llm"
	"tennis-coach-ai/internal/models"
)

type AnalysisService struct {
	llm *llm.Client
}

func (s *AnalysisService) Analyze(ctx context.Context, req models.AnalyzeRequest) (*models.AnalyzeResponse, error) {
	// TODO: build prompt
	// TODO: call LLM
	// TODO: parse response

	return nil, nil
}
