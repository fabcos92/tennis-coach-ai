package llm

import (
	"encoding/json"
	"errors"
	"tennis-coach-ai/internal/application/ports"
	"tennis-coach-ai/internal/domain/analysis"
)

type JSONMapper struct{}

func NewJSONMapper() ports.AnalysisMapper {
	return &JSONMapper{}
}

func (m *JSONMapper) FromLLM(raw string) (*analysis.Analysis, error) {
	var a analysis.Analysis

	if err := json.Unmarshal([]byte(raw), &a); err != nil {
		return nil, err
	}

	if a.Issues == nil {
		return nil, errors.New("invalid LLM response: issues is null")
	}

	return &a, nil
}
