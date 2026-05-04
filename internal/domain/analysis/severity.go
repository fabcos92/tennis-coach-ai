package analysis

type Severity string

const (
	Low    Severity = "low"
	Medium Severity = "medium"
	High   Severity = "high"
)

func (s Severity) String() string {
	return string(s)
}
