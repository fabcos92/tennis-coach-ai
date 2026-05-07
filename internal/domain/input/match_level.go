package input

import "strings"

type MatchLevel string

const (
	BeginnerMatchLevel MatchLevel = "beginner"
	MidMatchLevel      MatchLevel = "mid"
	ProMatchLevel      MatchLevel = "pro"
)

func NewMatchLevel(value string) (*MatchLevel, error) {
	MatchLevel := MatchLevel(strings.ToLower(value))

	if MatchLevel.IsValid() {
		return &MatchLevel, nil
	}

	return nil, ErrInvalidInput
}

func (l MatchLevel) String() string {
	return string(l)
}

func (l *MatchLevel) IsValid() bool {
	allowed := map[string]MatchLevel{
		"beginner": BeginnerMatchLevel,
		"mid":      MidMatchLevel,
		"pro":      ProMatchLevel,
	}

	_, ok := allowed[l.String()]

	return ok
}
