package dto

type AnalyzeResponse struct {
	Issues          []Issue  `json:"issues"`
	Recommendations []string `json:"recommendations"`
	FocusArea       string   `json:"focus_area"`
}

func NewAnalyzeResponse(issues []Issue, recommendations []string, focusArea string) AnalyzeResponse {
	return AnalyzeResponse{issues, recommendations, focusArea}
}
