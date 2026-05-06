package llm

import (
	"context"
	"log"
	"tennis-coach-ai/internal/application/ports"
	"time"
)

type Provider string

const (
	ProviderMock   Provider = "mock"
	ProviderOpenAI Provider = "openai"
	ProviderGroq   Provider = "groq"
)

type Gateway struct {
	policy    Policy
	providers []ProviderClient
}

func NewGateway(policy Policy, providers []ProviderClient) ports.LLM {
	return &Gateway{
		policy,
		providers,
	}
}

func (g *Gateway) Analyze(ctx context.Context, prompt string) (string, error) {
	var lastErr error

	for _, provider := range g.providers {
		if provider.Breaker != nil && !provider.Breaker.Allow() {
			log.Printf("[LLM] provider=%s skipped (circuit open)", provider.Name)
			continue
		}

		resp, err := g.callWithRetry(ctx, provider, prompt)
		if err == nil {
			if provider.Breaker != nil {
				provider.Breaker.Success()
			}
			return resp, nil
		}

		if provider.Breaker != nil {
			provider.Breaker.Fail()
		}

		lastErr = err
	}

	return "", lastErr
}

func (g *Gateway) callWithRetry(
	ctx context.Context,
	provider ProviderClient,
	prompt string,
) (string, error) {
	start := time.Now()
	ctx, cancel := context.WithTimeout(ctx, 6*time.Second)
	defer cancel()

	var lastErr error

	for i := 0; i < g.policy.MaxRetries; i++ {
		if ctx.Err() != nil {
			return "", ctx.Err()
		}

		resp, err := provider.LLM.Analyze(ctx, prompt)
		if err == nil {
			log.Printf("[LLM] provider=%s success duration=%s", provider.Name, time.Since(start))
			return resp, nil
		}

		log.Printf("[LLM] provider=%s failed err=%v", provider.Name, err)
		lastErr = err

		if !g.policy.Retryable(err) {
			return "", err
		}

		time.Sleep(g.policy.Backoff(i))
	}

	return "", lastErr
}
