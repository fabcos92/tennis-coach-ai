package commands

import (
	"context"
	"tennis-coach-ai/internal/application/ports"
	model "tennis-coach-ai/internal/domain/analysis"
)

type AnalyzeMatchPerformanceHandler struct {
	llm           ports.LLM
	mapper        ports.AnalysisMapper
	promptBuilder ports.PromptBuilder
}

func NewAnalyzeMatchPerformanceHandler(
	llm ports.LLM,
	mapper ports.AnalysisMapper,
	promptBuilder ports.PromptBuilder,
) *AnalyzeMatchPerformanceHandler {
	return &AnalyzeMatchPerformanceHandler{llm, mapper, promptBuilder}
}

func (h *AnalyzeMatchPerformanceHandler) Execute(ctx context.Context, command AnalyzeMatchPerformance) (*model.Analysis, error) {
	var prompt string
	if command.Type == "text" {
		input := ports.NewTextInput(command.TextInput.Text)
		prompt = h.promptBuilder.BuildText(input)
	} else {
		input := ports.NewStatsInput(command.StatsInput.FirstServeInPct, command.StatsInput.SecondServeInPct, command.StatsInput.UnforcedErrors)
		prompt = h.promptBuilder.BuildStats(input)
	}

	raw, err := h.llm.Analyze(ctx, prompt)
	if err != nil {
		return model.NewDefaultAnalysis(), err
	}

	analysis, err := h.mapper.FromLLM(raw)
	if err != nil {
		return model.NewDefaultAnalysis(), err
	}

	if err := analysis.Validate(); err != nil {
		return model.NewDefaultAnalysis(), err
	}

	return analysis, nil
}
