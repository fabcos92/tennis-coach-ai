package llm

import (
	"errors"
	"tennis-coach-ai/internal/infrastructure/llm/shared"
	"time"
)

type Policy struct {
	MaxRetries int
	Retryable  func(error) bool
	Backoff    func(attempt int) time.Duration
}

func DefaultPolicy() Policy {
	return Policy{
		MaxRetries: 3,
		Backoff: func(i int) time.Duration {
			return time.Duration(i+1) * 200 * time.Millisecond
		},
		Retryable: func(err error) bool {
			var llmErr shared.LLMError
			if errors.As(err, &llmErr) {
				return llmErr.Retryable
			}
			return false
		},
	}
}
