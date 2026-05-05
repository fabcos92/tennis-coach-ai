package dto

type AnalyzeRequest struct {
	Type string `json:"type"` // "match_stats" | "text"

	Stats MatchStats `json:"stats,omitempty"`
	Text  string     `json:"text,omitempty"`
}
