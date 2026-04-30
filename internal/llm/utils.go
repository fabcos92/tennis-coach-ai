package llm

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

func doRequest(ctx context.Context, client *http.Client, req *http.Request, provider string) (response, error) {
	resp, err := client.Do(req)
	if err != nil {
		return response{}, LLMError{
			Provider:  provider,
			Message:   err.Error(),
			Retryable: true,
		}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return response{}, LLMError{
			Provider:  provider,
			Message:   string(b),
			Retryable: resp.StatusCode == 429 || resp.StatusCode == 408 || resp.StatusCode >= 500,
		}
	}

	var r response
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return response{}, LLMError{
			Provider:  provider,
			Message:   err.Error(),
			Retryable: false,
		}
	}

	if len(r.Choices) == 0 {
		return response{}, LLMError{
			Provider:  provider,
			Message:   "empty response",
			Retryable: false,
		}
	}

	return r, nil
}

func withRetry(fn func() (string, error)) (string, error) {
	var lastErr error

	for i := 0; i < 3; i++ {
		res, err := fn()
		if err == nil {
			return res, nil
		}
		var llmErr LLMError
		if errors.As(err, &llmErr) {
			if !llmErr.Retryable {
				return "", err
			}
		}
		lastErr = err

		time.Sleep(time.Duration(200*(1<<i)) * time.Millisecond)
	}

	return "", lastErr
}
