package analysis

type Analysis struct {
	FocusArea       string
	Issues          []Issue
	Recommendations []string
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
