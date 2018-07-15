package main

import (
	"fmt"
)

// div is a <div> tag with a class.
type div struct {
	class string
	body  fmt.Stringer
}

func newDiv(class string, body fmt.Stringer) *div {
	return &div{
		class: class,
		body:  body,
	}
}

func (c *div) String() string {
	return fmt.Sprintf(`<div class="%s">%s</div>`, c.class, c.body)
}
