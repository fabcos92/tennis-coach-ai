package llm

import "time"

func withRetry(fn func() (string, error)) (string, error) {
	var lastErr error

	for i := 0; i < 2; i++ {
		res, err := fn()
		if err == nil {
			return res, nil
		}
		lastErr = err
		time.Sleep(200 * time.Millisecond)
	}

	return "", lastErr
}
