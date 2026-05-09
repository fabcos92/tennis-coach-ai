package analysis

type Issue struct {
	Text     string   `json:"text"`
	Severity Severity `json:"severity"`
}

func (i Issue) Validate() error {
	if i.Text == "" {
		return ErrEmptyIssueText
	}

	return i.Severity.Validate()
}
