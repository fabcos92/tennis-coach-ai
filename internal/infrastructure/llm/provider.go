package llm

import "tennis-coach-ai/internal/application/ports"

type ProviderClient struct {
	Name    string
	LLM     ports.LLM
	Breaker *CircuitBreaker
}
