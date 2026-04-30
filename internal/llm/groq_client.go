package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type GroqClient struct {
	apiKey     string
	httpClient *http.Client
	baseURL    string
}

func NewGroqClient(apiKey string) Client {
	return &GroqClient{
		apiKey:  apiKey,
		baseURL: "https://api.groq.com/openai/v1/chat/completions",
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *GroqClient) Analyze(ctx context.Context, prompt string) (string, error) {
	return withRetry(func() (string, error) {
		return c.callLLM(ctx, prompt)
	})
}

func (c *GroqClient) callLLM(ctx context.Context, prompt string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	reqBody := request{
		Model:       "llama-3.1-8b-instant",
		Temperature: 0.2,
		Messages: []message{
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
		return "", err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.baseURL,
		bytes.NewBuffer(bodyBytes),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var llmResp response
	json.NewDecoder(resp.Body).Decode(&llmResp)

	return llmResp.Choices[0].Message.Content, nil
}
