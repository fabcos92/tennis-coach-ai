package models

type Issue struct {
	Text     string `json:"text"`
	Severity string `json:"severity"`
}

type AnalyzeResponse struct {
	Issues          []Issue  `json:"issues"`
	Recommendations []string `json:"recommendations"`
	FocusArea       string   `json:"focus_area"`
}
