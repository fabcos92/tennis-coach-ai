package commands

import (
	"context"
	"tennis-coach-ai/internal/application/ports"
	model "tennis-coach-ai/internal/domain/analysis"
	"tennis-coach-ai/internal/domain/input"
)

type AnalyzeMatchPerformanceHandler struct {
	llm            ports.LLM
	analysisMapper ports.AnalysisMapper
	promptBuilder  ports.PromptBuilder
}

func NewAnalyzeMatchPerformanceHandler(
	llm ports.LLM,
	analysisMapper ports.AnalysisMapper,
	promptBuilder ports.PromptBuilder,
) *AnalyzeMatchPerformanceHandler {
	return &AnalyzeMatchPerformanceHandler{llm, analysisMapper, promptBuilder}
}

func (h *AnalyzeMatchPerformanceHandler) Execute(ctx context.Context, command AnalyzeMatchPerformance) (*model.Analysis, error) {
	inputType, err := h.toInputType(command.Type)
	if err != nil {
		return nil, err
	}

	var prompt string
	if inputType.IsText() {
		input, err := h.toText(command.Text)
		if err != nil {
			return nil, err
		}
		prompt = h.promptBuilder.BuildText(input)
	}
	if inputType.IsStats() {
		input, err := h.toStats(command.Stats)
		if err != nil {
			return nil, err
		}
		prompt = h.promptBuilder.BuildStats(input)
	}

	raw, err := h.llm.Analyze(ctx, prompt)
	if err != nil {
		return nil, err
	}

	analysis, err := h.analysisMapper.FromLLM(raw)
	if err != nil {
		return nil, err
	}

	if err := analysis.Validate(); err != nil {
		return nil, err
	}

	return analysis, nil
}

func (m *AnalyzeMatchPerformanceHandler) toStats(payload *StatsPayload) (*input.Stats, error) {
	firstServeInPct, err := input.NewPercent(payload.FirstServeInPct)
	if err != nil {
		return nil, err
	}
	firstServeWonPct, err := input.NewPercent(payload.FirstServeWonPct)
	if err != nil {
		return nil, err
	}
	firstServe := input.NewServe(firstServeInPct, firstServeWonPct)

	secondServeInPct, err := input.NewPercent(payload.SecondServeInPct)
	if err != nil {
		return nil, err
	}
	secondServeWonPct, err := input.NewPercent(payload.SecondServeWonPct)
	if err != nil {
		return nil, err
	}
	secondServe := input.NewServe(secondServeInPct, secondServeWonPct)

	serveReturnInPct, err := input.NewPercent(payload.ReturnInPct)
	if err != nil {
		return nil, err
	}
	serveReturnWonPct, err := input.NewPercent(payload.ReturnWonPct)
	if err != nil {
		return nil, err
	}
	serveReturn := input.NewServe(serveReturnInPct, serveReturnWonPct)

	surface, err := input.NewSurface(payload.Surface)
	if err != nil {
		return nil, err
	}

	matchLevel, err := input.NewMatchLevel(payload.MatchLevel)
	if err != nil {
		return nil, err
	}

	return input.NewStats(
		firstServe,
		secondServe,
		serveReturn,
		payload.Aces,
		payload.DoubleFaults,
		payload.Winners,
		payload.UnforcedErrors,
		surface,
		matchLevel,
	), nil
}

func (m *AnalyzeMatchPerformanceHandler) toText(text string) (*input.Text, error) {
	return input.NewText(text), nil
}

func (m *AnalyzeMatchPerformanceHandler) toInputType(inputType string) (*input.InputType, error) {
	return input.NewInputType(inputType)
}
