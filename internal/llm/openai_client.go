package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
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
		apiKey:     apiKey,
		baseURL:    "https://api.openai.com/v1/chat/completions",
		httpClient: &http.Client{},
	}
}

func (c *OpenAIClient) Analyze(ctx context.Context, prompt string) (string, error) {
	return withRetry(func() (string, error) {
		return c.callLLM(ctx, prompt)
	})
}

func (c *OpenAIClient) callLLM(ctx context.Context, prompt string) (string, error) {
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

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("llm error: %s", string(b))
	}

	var llmResp response

	err = json.NewDecoder(resp.Body).Decode(&llmResp)
	if err != nil {
		return "", err
	}

	if len(llmResp.Choices) == 0 {
		return "", fmt.Errorf("empty response")
	}

	return llmResp.Choices[0].Message.Content, nil

}
