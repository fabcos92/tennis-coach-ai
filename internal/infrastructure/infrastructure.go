package infrastructure

import (
	config "tennis-coach-ai/cfg"
	"tennis-coach-ai/internal/application/ports"
	"tennis-coach-ai/internal/infrastructure/llm"
	"tennis-coach-ai/internal/infrastructure/logging"
	"time"
)

type Infrastructure struct {
	LLM           ports.LLM
	Mapper        ports.AnalysisMapper
	PromptBuilder ports.PromptBuilder
}

func New(cfg *config.Config) *Infrastructure {
	logger := logging.NewStdLogger()
	openai := llm.NewOpenAI(cfg)
	groq := llm.NewGroq(cfg)
	mock := llm.NewMock()

	gateway := llm.NewGateway(
		llm.DefaultPolicy(),
		[]llm.ProviderClient{
			{
				Name:    "groq",
				LLM:     llm.NewLoggingLLM(groq, "groq", logger),
				Breaker: llm.NewCircuitBreaker(3, 30*time.Second),
			},
			{
				Name:    "openai",
				LLM:     llm.NewLoggingLLM(openai, "openai", logger),
				Breaker: llm.NewCircuitBreaker(3, 30*time.Second),
			},
			{Name: "mock", LLM: mock},
		},
	)

	return &Infrastructure{
		LLM:           gateway,
		Mapper:        llm.NewJSONMapper(),
		PromptBuilder: llm.NewDefaultPromptBuilder(),
	}
}
