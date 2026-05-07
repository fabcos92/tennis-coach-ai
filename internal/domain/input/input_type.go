package input

import "strings"

type InputType string

const (
	StatsInputType InputType = "stats"
	TextInputType  InputType = "text"
	VideoInputType InputType = "video"
)

func NewInputType(value string) (*InputType, error) {
	inputType := InputType(strings.ToLower(value))

	if inputType.IsValid() {
		return &inputType, nil
	}

	return nil, ErrInvalidInput
}

func (t InputType) String() string {
	return string(t)
}

func (t *InputType) IsValid() bool {
	allowed := map[string]InputType{
		"stats": StatsInputType,
		"text":  TextInputType,
		"video": VideoInputType,
	}

	_, ok := allowed[t.String()]

	return ok
}

func (t InputType) IsStats() bool {
	return t == StatsInputType
}

func (t InputType) IsText() bool {
	return t == TextInputType
}
