package models

type AnalyzeResponse struct {
	Issues          []string `json:"issues"`
	Recommendations []string `json:"recommendations"`
	FocusArea       string   `json:"focus_area"`
}
