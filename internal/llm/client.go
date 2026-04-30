package llm

import "context"

type Client interface {
	Analyze(ctx context.Context, prompt string) (string, error)
}
