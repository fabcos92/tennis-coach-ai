package input

import "strings"

type Surface string

const (
	ClaySurface  Surface = "clay"
	HardSurface  Surface = "hard"
	GrassSurface Surface = "grass"
)

func NewSurface(value string) (*Surface, error) {
	surface := Surface(strings.ToLower(value))

	if surface.IsValid() {
		return &surface, nil
	}

	return nil, ErrInvalidInput
}

func (s Surface) String() string {
	return string(s)
}

func (s *Surface) IsValid() bool {
	allowed := map[string]Surface{
		"clay":  ClaySurface,
		"hard":  HardSurface,
		"grass": GrassSurface,
	}

	_, ok := allowed[s.String()]

	return ok
}
