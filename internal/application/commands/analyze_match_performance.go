package commands

type AnalyzeMatchPerformance struct {
	Type       string
	StatsInput struct {
		FirstServeInPct  int
		SecondServeInPct int
		UnforcedErrors   int
	}
	TextInput struct {
		Text string
	}
}

func NewAnalyzeMatchPerformance(analysisType, text string, firstServeInPct, secondServeInPct, unforcedErrors int) AnalyzeMatchPerformance {
	command := AnalyzeMatchPerformance{}
	command.Type = analysisType
	command.StatsInput.FirstServeInPct = firstServeInPct
	command.StatsInput.SecondServeInPct = secondServeInPct
	command.StatsInput.UnforcedErrors = unforcedErrors

	return command
}
