package html

// Text allows text to be rendered on the page.
type Text struct {
	s string
}

func NewText(s string) *Text {
	return &Text{
		s: s,
	}
}

func (c *Text) String() string {
	return c.s
}
