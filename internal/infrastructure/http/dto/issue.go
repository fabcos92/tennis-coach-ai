package dto

type Issue struct {
	Text     string `json:"text"`
	Severity string `json:"severity"`
}

func NewIssue(text, severity string) Issue {
	return Issue{text, severity}
}
