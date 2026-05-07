package input

type Percent float64

func NewPercent(value float64) (*Percent, error) {
	if value < 0 && value > 100 {
		return nil, ErrInvalidInput
	}

	percent := Percent(value)
	return &percent, nil
}
