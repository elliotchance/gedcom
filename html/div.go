package html

import (
	"fmt"
)

// Div is a <div> tag with a class.
type Div struct {
	class string
	body  fmt.Stringer
}

func NewDiv(class string, body fmt.Stringer) *Div {
	return &Div{
		class: class,
		body:  body,
	}
}

func (c *Div) String() string {
	return fmt.Sprintf(`<div class="%s">%s</div>`, c.class, c.body)
}
