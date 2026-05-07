package dto

type MatchStats struct {
	FirstServeInPct  float64 `json:"first_serve_in_pct"`
	SecondServeInPct float64 `json:"second_serve_in_pct"`

	FirstServeWonPct  float64 `json:"first_serve_won_pct"`
	SecondServeWonPct float64 `json:"second_serve_won_pct"`

	ReturnInPct  float64 `json:"return_in_pct"`
	ReturnWonPct float64 `json:"return_won_pct"`

	Aces         int `json:"aces"`
	DoubleFaults int `json:"double_faults"`

	Winners        int `json:"winners"`
	UnforcedErrors int `json:"unforced_errors"`

	Surface    string `json:"surface"`
	MatchLevel string `json:"match_level"`
}
