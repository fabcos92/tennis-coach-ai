package commands

import "tennis-coach-ai/internal/domain/input"

type AnalyzeMatchPerformance struct {
	Type       *input.InputType
	StatsInput *input.Stats
	TextInput  *input.Text
}

func NewAnalyzeMatchPerformance(inputType *input.InputType, statsInput *input.Stats, textInput *input.Text) AnalyzeMatchPerformance {
	command := AnalyzeMatchPerformance{
		Type:       inputType,
		StatsInput: statsInput,
		TextInput:  textInput,
	}

	return command
}
