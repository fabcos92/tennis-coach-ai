package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	config "tennis-coach-ai/cfg"
	"tennis-coach-ai/internal/application/ports"
	"tennis-coach-ai/internal/infrastructure/llm/shared"
	"time"
)

type Groq struct {
	apiKey     string
	httpClient *http.Client
	baseURL    string
}

func NewGroq(cfg *config.Config) ports.LLM {
	return &Groq{
		apiKey:  cfg.Groq.Key,
		baseURL: "https://api.groq.com/openai/v1/chat/completions",
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Groq) Analyze(ctx context.Context, prompt string) (string, error) {
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
		return "", shared.LLMError{
			Provider:  c.name(),
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
		return "", shared.LLMError{
			Provider:  c.name(),
			Message:   string(err.Error()),
			Retryable: false,
		}
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := shared.DoRequest(ctx, c.httpClient, req, c.name())
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func (c *Groq) name() string {
	return "Groq"
}
