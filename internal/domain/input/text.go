package input

type Text struct {
	Text string
}

func NewText(text string) *Text {
	return &Text{text}
}
