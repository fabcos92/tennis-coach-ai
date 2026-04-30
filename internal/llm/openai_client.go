package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type OpenAIClient struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

func NewOpenAIClient(apiKey string) Client {
	return &OpenAIClient{
		apiKey:  apiKey,
		baseURL: "https://api.openai.com/v1/chat/completions",
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *OpenAIClient) Analyze(ctx context.Context, prompt string) (string, error) {
	return withRetry(func() (string, error) {
		return c.callLLM(ctx, prompt)
	})
}

func (c *OpenAIClient) callLLM(ctx context.Context, prompt string) (string, error) {
	log.Printf("[LLM] calling provider (model=%s)", "gpt-4o-mini")
	start := time.Now()
	defer func() {
		log.Printf("[LLM] request finished in %s", time.Since(start))
	}()

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	reqBody := request{
		Model:       "gpt-4o-mini",
		Temperature: 0.2,
		Messages: []message{
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
		log.Printf("[LLM] request failed: %v", err)
		return "", LLMError{
			Provider:  "openai",
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
		return "", LLMError{
			Provider:  "openai",
			Message:   string(err.Error()),
			Retryable: false,
		}
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := doRequest(ctx, c.httpClient, req, "openai")
	if err != nil {
		log.Printf("[LLM] request failed: %v", err)
		return "", err
	}
	log.Printf("[LLM] response received successfully")
	return resp.Choices[0].Message.Content, nil
}
