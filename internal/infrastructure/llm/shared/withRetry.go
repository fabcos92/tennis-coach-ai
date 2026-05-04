package shared

import (
	"errors"
	"time"
)

func WithRetry(fn func() (string, error)) (string, error) {
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
