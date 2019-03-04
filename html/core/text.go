package core

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

func (c *Text) WriteHTMLTo(w io.Writer) (int64, error) {
	s := strings.Replace(c.s, "&nbsp;", "~~space~~", -1)
	s = html.EscapeString(s)
	s = strings.Replace(s, "~~space~~", "&nbsp;", -1)

	return writeString(w, s)
}
