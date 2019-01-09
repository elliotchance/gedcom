package html

import (
	"html"
	"io"
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

func (c *Text) WriteTo(w io.Writer) (int64, error) {
	s := strings.Replace(c.s, "&nbsp;", "~~space~~", -1)
	s = html.EscapeString(s)

	return writeString(w, strings.Replace(s, "~~space~~", "&nbsp;", -1))
}
