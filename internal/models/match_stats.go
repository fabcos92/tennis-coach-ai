package models

type MatchStats struct {
	FirstServePct     int `json:"first_serve_pct"`
	SecondServeWonPct int `json:"second_serve_won_pct"`
	UnforcedErrors    int `json:"unforced_errors"`
}
