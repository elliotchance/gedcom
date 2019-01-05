package html

import (
	"html"
	"strings"
)

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
	s := strings.Replace(c.s, "&nbsp;", "~~space~~", -1)
	s = html.EscapeString(s)

	return strings.Replace(s, "~~space~~", "&nbsp;", -1)
}
