package shared

import "fmt"

type LLMError struct {
	Provider  string
	Message   string
	Retryable bool
}

func (e LLMError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Provider, e.Message)
}
