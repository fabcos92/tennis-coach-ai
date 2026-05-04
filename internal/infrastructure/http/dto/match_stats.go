package dto

type MatchStats struct {
	FirstServeInPct  int `json:"first_serve_in_pct"`
	SecondServeInPct int `json:"second_serve_in_pct"`
	UnforcedErrors   int `json:"unforced_errors"`
}
