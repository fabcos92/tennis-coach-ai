package shared

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

func DoRequest(ctx context.Context, client *http.Client, req *http.Request, provider string) (*Response, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, LLMError{
			Provider:  provider,
			Message:   err.Error(),
			Retryable: true,
		}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return nil, LLMError{
			Provider:  provider,
			Message:   string(b),
			Retryable: resp.StatusCode == 429 || resp.StatusCode == 408 || resp.StatusCode >= 500,
		}
	}

	r := &Response{}
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, LLMError{
			Provider:  provider,
			Message:   err.Error(),
			Retryable: false,
		}
	}

	if len(r.Choices) == 0 {
		return nil, LLMError{
			Provider:  provider,
			Message:   "empty response",
			Retryable: false,
		}
	}

	return r, nil
}
