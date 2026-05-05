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

type OpenAI struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

func NewOpenAI(cfg *config.Config) ports.LLM {
	return &OpenAI{
		apiKey:  cfg.OpenAI.Key,
		baseURL: "https://api.openai.com/v1/chat/completions",
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *OpenAI) Analyze(ctx context.Context, prompt string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	reqBody := shared.Request{
		Model:       "gpt-4o-mini",
		Temperature: 0.2,
		Messages: []shared.Message{
			{
				Role:    "system",
				Content: "You are a precise assistant that outputs only valid JSON.",
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

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := shared.DoRequest(ctx, c.httpClient, req, c.name())
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}

func (c *OpenAI) name() string {
	return "Open AI"
}
