package mappers

import (
	"tennis-coach-ai/internal/domain/input"
	"tennis-coach-ai/internal/infrastructure/http/dto"
)

type MatchInputMapper interface {
	ToStats(req dto.MatchStats) (*input.Stats, error)
	ToText(text string) (*input.Text, error)
	ToInputType(inputType string) (*input.InputType, error)
}

type DefaultMatchInputMapper struct {
}

func NewDefaultMatchInputMapper() MatchInputMapper {
	return &DefaultMatchInputMapper{}
}

func (m *DefaultMatchInputMapper) ToStats(req dto.MatchStats) (*input.Stats, error) {
	firstServeInPct, err := input.NewPercent(req.FirstServeInPct)
	if err != nil {
		return nil, err
	}
	firstServeWonPct, err := input.NewPercent(req.FirstServeWonPct)
	if err != nil {
		return nil, err
	}
	firstServe := input.NewServe(firstServeInPct, firstServeWonPct)

	secondServeInPct, err := input.NewPercent(req.SecondServeInPct)
	if err != nil {
		return nil, err
	}
	secondServeWonPct, err := input.NewPercent(req.SecondServeWonPct)
	if err != nil {
		return nil, err
	}
	secondServe := input.NewServe(secondServeInPct, secondServeWonPct)

	serveReturnInPct, err := input.NewPercent(req.ReturnInPct)
	if err != nil {
		return nil, err
	}
	serveReturnWonPct, err := input.NewPercent(req.ReturnWonPct)
	if err != nil {
		return nil, err
	}
	serveReturn := input.NewServe(serveReturnInPct, serveReturnWonPct)

	surface, err := input.NewSurface(req.Surface)
	if err != nil {
		return nil, err
	}

	matchLevel, err := input.NewMatchLevel(req.MatchLevel)
	if err != nil {
		return nil, err
	}

	return input.NewStats(
		firstServe,
		secondServe,
		serveReturn,
		req.Aces,
		req.DoubleFaults,
		req.Winners,
		req.UnforcedErrors,
		surface,
		matchLevel,
	), nil
}

func (m *DefaultMatchInputMapper) ToText(text string) (*input.Text, error) {
	return input.NewText(text), nil
}

func (m *DefaultMatchInputMapper) ToInputType(inputType string) (*input.InputType, error) {
	return input.NewInputType(inputType)
}
