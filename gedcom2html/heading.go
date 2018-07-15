package main

import (
	"fmt"
)

// heading is larger text.
type heading struct {
	text, class string
	number      int
}

func newHeading(number int, class, text string) *heading {
	return &heading{
		text:   text,
		number: number,
		class:  class,
	}
}

func (c *heading) String() string {
	return fmt.Sprintf(`<h%d class="%s">%s</h%d>`,
		c.number, c.class, c.text, c.number)
}
