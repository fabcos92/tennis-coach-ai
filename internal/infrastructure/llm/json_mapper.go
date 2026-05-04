package llm

import (
	"encoding/json"
	"tennis-coach-ai/internal/application/ports"
	"tennis-coach-ai/internal/domain/analysis"
)

type JSONMapper struct{}

func NewJSONMapper() ports.AnalysisMapper {
	return &JSONMapper{}
}

func (m *JSONMapper) FromLLM(raw string) (*analysis.Analysis, error) {
	var a *analysis.Analysis
	err := json.Unmarshal([]byte(raw), &a)
	return a, err
}
