package ports

import "context"

type LLM interface {
	Analyze(ctx context.Context, prompt string) (string, error)
}
