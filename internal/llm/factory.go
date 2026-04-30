package llm

type Provider string

const (
	ProviderMock   Provider = "mock"
	ProviderOpenAI Provider = "openai"
)

func NewProvider(provider Provider, apiKey string) Client {
	switch provider {

	case ProviderOpenAI:
		return NewOpenAIClient(apiKey)

	case ProviderMock:
		fallthrough
	default:
		return NewMockClient()
	}
}
