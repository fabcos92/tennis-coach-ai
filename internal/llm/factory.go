package llm

type Provider string

const (
	ProviderMock   Provider = "mock"
	ProviderOpenAI Provider = "openai"
	ProviderGroq   Provider = "groq"
)

func NewProvider(provider Provider, openAIApiKey, groqAPIKey string) Client {
	switch provider {

	case ProviderOpenAI:
		return NewOpenAIClient(openAIApiKey)

	case ProviderGroq:
		return NewGroqClient(groqAPIKey)

	case ProviderMock:
		fallthrough
	default:
		return NewMockClient()
	}
}
