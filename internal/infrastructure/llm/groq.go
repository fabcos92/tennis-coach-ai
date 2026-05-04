package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	config "tennis-coach-ai/cfg"
	"tennis-coach-ai/internal/application/ports"
	"tennis-coach-ai/internal/infrastructure/llm/shared"
	"time"
)

type Groq struct {
	apiKey       string
	httpClient   *http.Client
	baseURL      string
	providerName string
}

func NewGroq(cfg *config.Config) ports.LLM {
	return &Groq{
		apiKey:  cfg.Groq.Key,
		baseURL: "https://api.groq.com/openai/v1/chat/completions",
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		providerName: "groq",
	}
}

func (c *Groq) Analyze(ctx context.Context, prompt string) (string, error) {
	return shared.WithRetry(func() (string, error) {
		return c.callLLM(ctx, prompt)
	})
}

func (c *Groq) callLLM(ctx context.Context, prompt string) (string, error) {
	log.Printf("[LLM] calling provider (model=%s)", "llama-3.1-8b-instant")
	start := time.Now()
	defer func() {
		log.Printf("[LLM] request finished in %s", time.Since(start))
	}()

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	reqBody := shared.Request{
		Model:       "llama-3.1-8b-instant",
		Temperature: 0.2,
		Messages: []shared.Message{
			{
				Role:    "system",
				Content: "You must return ONLY valid JSON.",
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		log.Printf("[LLM] request failed: %v", err)
		return "", shared.LLMError{
			Provider:  c.providerName,
			Message:   string(err.Error()),
			Retryable: false,
		}
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.baseURL,
		bytes.NewBuffer(bodyBytes),
	)
	if err != nil {
		log.Printf("[LLM] request failed: %v", err)
		return "", shared.LLMError{
			Provider:  c.providerName,
			Message:   string(err.Error()),
			Retryable: false,
		}
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := shared.DoRequest(ctx, c.httpClient, req, c.providerName)
	if err != nil {
		log.Printf("[LLM] request failed: %v", err)
		return "", err
	}
	log.Printf("[LLM] response received successfully")
	return resp.Choices[0].Message.Content, nil
}
