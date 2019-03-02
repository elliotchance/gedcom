package core

import "io"

type HTML struct {
	s string
}

func NewHTML(s string) *HTML {
	return &HTML{
		s: s,
	}
}

func (c *HTML) WriteHTMLTo(w io.Writer) (int64, error) {
	return writeString(w, c.s)
}
