package analysis

import "strings"

type FocusArea string

const (
	Serve       FocusArea = "serve"
	Return      FocusArea = "return"
	Consistency FocusArea = "consistency"
	Movement    FocusArea = "movement"
	Aggression  FocusArea = "aggression"
	NetPlay     FocusArea = "net_play"
	Mental      FocusArea = "mental"
	Fitness     FocusArea = "fitness"
)

func (s FocusArea) String() string {
	return string(s)
}

func (f FocusArea) Normalize() FocusArea {
	switch strings.ToLower(string(f)) {
	case "netplay", "net-play":
		return NetPlay
	default:
		return FocusArea(strings.ToLower(string(f)))
	}
}

func (s FocusArea) Validate() error {
	allowed := map[string]FocusArea{
		"serve":       Serve,
		"return":      Return,
		"consistency": Consistency,
		"movement":    Movement,
		"aggression":  Aggression,
		"net_play":    NetPlay,
		"mental":      Mental,
		"fitness":     Fitness,
	}

	_, ok := allowed[s.String()]

	if !ok {
		return ErrInvalidFocusArea
	}

	return nil
}
