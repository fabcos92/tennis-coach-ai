package analysis

import "strings"

type Severity string

const (
	Low    Severity = "low"
	Medium Severity = "medium"
	High   Severity = "high"
)

func (s Severity) String() string {
	return string(s)
}

func (s Severity) Normalize() Severity {
	return Severity(strings.ToLower(string(s)))
}

func (s Severity) Validate() error {
	allowed := map[string]Severity{
		"low":    Low,
		"medium": Medium,
		"high":   High,
	}

	_, ok := allowed[s.String()]

	if !ok {
		return ErrInvalidSeverity
	}

	return nil
}
