package llm

import (
	"context"
	"tennis-coach-ai/internal/application/ports"
)

type Mock struct{}

func NewMock() ports.LLM {
	return &Mock{}
}

func (m *Mock) Analyze(ctx context.Context, prompt string) (string, error) {
	return `{
		"issues": [{"text": "Weak second serve", "severity": "medium"}],
		"recommendations": ["Practice second serve consistency"],
		"focus_area": "serve"
	}`, nil
}
