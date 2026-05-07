package commands

type StatsPayload struct {
	FirstServeInPct  float64
	SecondServeInPct float64

	FirstServeWonPct  float64
	SecondServeWonPct float64

	ReturnInPct  float64
	ReturnWonPct float64

	Aces         int
	DoubleFaults int

	Winners        int
	UnforcedErrors int

	Surface    string
	MatchLevel string
}

func NewStatsPayload(
	firstServeInPct, firstServeWonPct, secondServeInPct, secondServeWonPct, returnInPct, returnWonPct float64,
	aces, doubleFaults, winners, unforcedErrors int,
	surface, matchLevel string,
) *StatsPayload {
	return &StatsPayload{
		FirstServeInPct:   firstServeInPct,
		SecondServeInPct:  firstServeWonPct,
		FirstServeWonPct:  secondServeInPct,
		SecondServeWonPct: secondServeWonPct,
		ReturnInPct:       returnInPct,
		ReturnWonPct:      returnWonPct,
		Aces:              aces,
		DoubleFaults:      doubleFaults,
		Winners:           winners,
		UnforcedErrors:    unforcedErrors,
		Surface:           surface,
		MatchLevel:        matchLevel,
	}
}

type AnalyzeMatchPerformance struct {
	Type string

	Stats *StatsPayload
	Text  string
}

func NewAnalyzeMatchPerformance(inputType string, statsInput *StatsPayload, textInput string) AnalyzeMatchPerformance {
	command := AnalyzeMatchPerformance{
		Type:  inputType,
		Stats: statsInput,
		Text:  textInput,
	}

	return command
}
