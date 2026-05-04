package ports

type StatsInput struct {
	FirstServeInPct  int
	SecondServeInPct int
	UnforcedErrors   int
}

func NewStatsInput(firstServeInPct, secondServeInPct, unforcedErrors int) StatsInput {
	return StatsInput{firstServeInPct, secondServeInPct, unforcedErrors}
}

type TextInput struct {
	Text string
}

func NewTextInput(text string) TextInput {
	return TextInput{text}
}

type PromptBuilder interface {
	BuildStats(input StatsInput) string
	BuildText(input TextInput) string
}
