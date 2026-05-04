package ports

import "tennis-coach-ai/internal/domain/analysis"

type AnalysisMapper interface {
	FromLLM(raw string) (*analysis.Analysis, error)
}
