package dto

type AnalyzeRequest struct {
	Type string `json:"type"` // "stats" | "text" | "video"

	Stats MatchStats `json:"stats,omitempty"`
	Text  string     `json:"text,omitempty"`
}
