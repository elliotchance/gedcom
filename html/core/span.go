package core

import (
	"io"
)

type Span struct {
	class string
	value Component
}

func NewSpan(class string, value Component) *Span {
	return &Span{
		value: value,
		class: class,
	}
}

func (c *Span) WriteHTMLTo(w io.Writer) (int64, error) {
	return NewTag("span", map[string]string{
		"class": c.class,
	}, c.value).WriteHTMLTo(w)
}
