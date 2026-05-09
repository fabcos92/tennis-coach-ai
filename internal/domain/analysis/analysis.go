package analysis

type Analysis struct {
	FocusArea       FocusArea `json:"focus_area"`
	Issues          []Issue   `json:"issues"`
	Recommendations []string  `json:"recommendations"`
}

func NewDefaultAnalysis() *Analysis {
	return &Analysis{
		FocusArea: "unknown",
	}
}

func (a *Analysis) Normalize() {
	for i := range a.Issues {
		a.Issues[i].Severity = a.Issues[i].Severity.Normalize()
	}

	a.FocusArea = a.FocusArea.Normalize()
}

func (a *Analysis) Validate() error {
	if err := a.FocusArea.Validate(); err != nil {
		return err
	}

	if len(a.Recommendations) > 5 {
		return ErrRecommendationsSizeExceeded
	}

	if len(a.Issues) > 3 {
		return ErrIssuesSizeExceeded
	}

	for _, issue := range a.Issues {
		if err := issue.Validate(); err != nil {
			return err
		}
	}

	return nil
}
