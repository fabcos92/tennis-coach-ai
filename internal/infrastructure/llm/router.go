package llm

import (
	"context"
	"tennis-coach-ai/internal/application/ports"
)

type Provider string

const (
	ProviderMock   Provider = "mock"
	ProviderOpenAI Provider = "openai"
	ProviderGroq   Provider = "groq"
)

type Router struct {
	provider Provider
	clients  map[Provider]ports.LLM
	// strategy RoutingStrategy // future
}

func NewRouter(
	provider Provider,
	groq ports.LLM,
	openai ports.LLM,
	mock ports.LLM,
) ports.LLM {
	return &Router{
		provider: provider,
		clients: map[Provider]ports.LLM{
			ProviderGroq:   groq,
			ProviderOpenAI: openai,
			ProviderMock:   mock,
		},
	}
}

func (r *Router) Analyze(ctx context.Context, prompt string) (string, error) {
	client, ok := r.clients[r.provider]
	if !ok {
		client = r.clients[ProviderMock]
	}
	return client.Analyze(ctx, prompt)
}
