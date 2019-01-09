package html

import "io"

type HTML struct {
	s string
}

func NewHTML(s string) *HTML {
	return &HTML{
		s: s,
	}
}

func (c *HTML) WriteTo(w io.Writer) (int64, error) {
	return writeString(w, c.s)
}
