package llm

type Client struct {
	apiKey string
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
	}
}

func (c *Client) Analyze(prompt string) (string, error) {
	// mock
	return `{
		"issues": ["Weak second serve"],
		"recommendations": ["Train second serve consistency drills"],
		"focus_area": "serve"
	}`, nil
}
