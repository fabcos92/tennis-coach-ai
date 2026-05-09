package llm

import (
	"encoding/json"
	"strings"
	"tennis-coach-ai/internal/application/ports"
	"tennis-coach-ai/internal/domain/analysis"
)

type JSONMapper struct{}

func NewJSONMapper() ports.AnalysisMapper {
	return &JSONMapper{}
}

func (m *JSONMapper) FromLLM(raw string) (*analysis.Analysis, error) {
	var a analysis.Analysis

	clean := sanitizeJSON(raw)

	if err := json.Unmarshal([]byte(clean), &a); err != nil {
		return nil, err
	}

	a.Normalize()

	if err := a.Validate(); err != nil {
		return nil, err
	}

	return &a, nil
}

func sanitizeJSON(raw string) string {
	raw = strings.TrimSpace(raw)
	raw = strings.Trim(raw, "`")
	start := strings.Index(raw, "{")
	end := strings.LastIndex(raw, "}")

	if start == -1 || end == -1 || start > end {
		return raw
	}

	return raw[start : end+1]
}
