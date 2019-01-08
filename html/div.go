package html

import (
	"io"
)

// Div is a <div> tag with a class.
type Div struct {
	class string
	body  Component
}

func NewDiv(class string, body Component) *Div {
	return &Div{
		class: class,
		body:  body,
	}
}

func (c *Div) WriteTo(w io.Writer) (int64, error) {
	attributes := map[string]string{
		"class": c.class,
	}

	return NewTag("div", attributes, c.body).WriteTo(w)
}
