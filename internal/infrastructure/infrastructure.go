package infrastructure

import (
	config "tennis-coach-ai/cfg"
	"tennis-coach-ai/internal/application/ports"
	"tennis-coach-ai/internal/infrastructure/llm"
)

type Infrastructure struct {
	LLM           ports.LLM
	Mapper        ports.AnalysisMapper
	PromptBuilder ports.PromptBuilder
}

func New(cfg *config.Config) *Infrastructure {
	openai := llm.NewOpenAI(cfg)
	groq := llm.NewGroq(cfg)
	mock := llm.NewMock()

	router := llm.NewRouter(llm.Provider(cfg.LLM.Provider), openai, groq, mock)

	return &Infrastructure{
		LLM:           router,
		Mapper:        llm.NewJSONMapper(),
		PromptBuilder: llm.NewDefaultPromptBuilder(),
	}
}
