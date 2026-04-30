package llm

import (
	"context"
)

type MockClient struct{}

func NewMockClient() Client {
	return &MockClient{}
}

func (m *MockClient) Analyze(ctx context.Context, prompt string) (string, error) {
	return `{
		"issues": ["Weak second serve"],
		"recommendations": ["Practice second serve consistency"],
		"focus_area": "serve"
	}`, nil
}
