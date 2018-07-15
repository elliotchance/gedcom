package main

// text allows text to be rendered on the page.
type text struct {
	s string
}

func newText(s string) *text {
	return &text{
		s: s,
	}
}

func (c *text) String() string {
	return c.s
}
