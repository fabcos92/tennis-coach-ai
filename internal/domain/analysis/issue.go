package analysis

type Issue struct {
	Text     string   `json:"text"`
	Severity Severity `json:"severity"`
}
