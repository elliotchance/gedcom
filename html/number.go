package html

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type number struct {
	value int
}

func NewNumber(value int) *number {
	return &number{
		value: value,
	}
}

func (c *number) String() string {
	p := message.NewPrinter(language.English)

	return p.Sprintf("%d", c.value)
}
