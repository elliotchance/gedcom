package core

import (
	"fmt"
	"io"
)

// BadgePill is a rounded badge that contains a value.
type BadgePill struct {
	color, class string
	value        Component
}

func NewBadgePill(color, class string, value Component) *BadgePill {
	return &BadgePill{
		color: color,
		value: value,
		class: class,
	}
}

func (c *BadgePill) WriteHTMLTo(w io.Writer) (int64, error) {
	class := fmt.Sprintf("badge badge-pill badge-%s %s", c.color, c.class)

	return NewSpan(class, c.value).WriteHTMLTo(w)
}
