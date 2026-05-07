package mappers

import (
	"tennis-coach-ai/internal/application/commands"
	"tennis-coach-ai/internal/infrastructure/http/dto"
)

type AnalyzeRequestMapper struct{}

func NewAnalyzeRequestMapper() *AnalyzeRequestMapper {
	return &AnalyzeRequestMapper{}
}

func (m *AnalyzeRequestMapper) ToCommand(req dto.AnalyzeRequest) commands.AnalyzeMatchPerformance {
	var stats *commands.StatsPayload
	if req.Stats != nil {
		stats = commands.NewStatsPayload(
			req.Stats.FirstServeInPct,
			req.Stats.FirstServeWonPct,
			req.Stats.SecondServeInPct,
			req.Stats.SecondServeWonPct,
			req.Stats.ReturnInPct,
			req.Stats.ReturnWonPct,
			req.Stats.Aces,
			req.Stats.DoubleFaults,
			req.Stats.Winners,
			req.Stats.UnforcedErrors,
			req.Stats.Surface,
			req.Stats.MatchLevel,
		)
	}

	return commands.NewAnalyzeMatchPerformance(
		req.Type,
		stats,
		req.Text,
	)
}
