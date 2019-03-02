package core

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"io"
)

type Number struct {
	value int
}

func NewNumber(value int) *Number {
	return &Number{
		value: value,
	}
}

func (c *Number) WriteHTMLTo(w io.Writer) (int64, error) {
	p := message.NewPrinter(language.English)

	return writeString(w, p.Sprintf("%d", c.value))
}
