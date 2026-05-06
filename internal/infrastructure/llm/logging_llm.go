package llm

import (
	"context"
	"tennis-coach-ai/internal/application/ports"
	"time"
)

type LoggingLLM struct {
	next   ports.LLM
	name   string
	logger ports.Logger
}

func NewLoggingLLM(next ports.LLM, name string, logger ports.Logger) ports.LLM {
	return &LoggingLLM{next: next, name: name, logger: logger}
}

func (l *LoggingLLM) Analyze(ctx context.Context, prompt string) (string, error) {
	start := time.Now()

	resp, err := l.next.Analyze(ctx, prompt)

	duration := time.Since(start)

	if err != nil {
		l.logger.Error("llm_call_failed",
			"provider", l.name,
			"duration", duration,
			"err", err,
		)
		return "", err
	}

	l.logger.Info("llm_call_success",
		"provider", l.name,
		"duration", duration,
	)

	return resp, nil
}
