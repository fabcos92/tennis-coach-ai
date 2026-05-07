package ports

import "tennis-coach-ai/internal/domain/input"

type PromptBuilder interface {
	BuildStats(i *input.Stats) string
	BuildText(i *input.Text) string
}
