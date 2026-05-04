package application

import (
	"tennis-coach-ai/internal/application/commands"
	"tennis-coach-ai/internal/infrastructure"
)

type Application struct {
	Commands struct {
		AnalyzeMatchPerformance *commands.AnalyzeMatchPerformanceHandler
	}
}

func NewApplication(infra *infrastructure.Infrastructure) *Application {
	application := &Application{}

	application.Commands.AnalyzeMatchPerformance = commands.NewAnalyzeMatchPerformanceHandler(infra.LLM, infra.Mapper, infra.PromptBuilder)

	return application
}
