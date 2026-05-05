package analysis

type Analysis struct {
	FocusArea       string   `json:"focus_area"`
	Issues          []Issue  `json:"issues"`
	Recommendations []string `json:"recommendations"`
}

func NewDefaultAnalysis() *Analysis {
	return &Analysis{
		FocusArea: "unknown",
	}
}

func (a *Analysis) Validate() error {
	if a.FocusArea == "" {
		return ErrInvalidAnalysis
	}

	return nil
}
