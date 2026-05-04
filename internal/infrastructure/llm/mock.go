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
		"issues": ["Weak second serve"],
		"recommendations": ["Practice second serve consistency"],
		"focus_area": "serve"
	}`, nil
}
