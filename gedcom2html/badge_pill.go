package main

import (
	"fmt"
)

// badgePill is a rounded badge that contains a value.
type badgePill struct {
	color, class, value string
}

func newBadgePill(color, class, value string) *badgePill {
	return &badgePill{
		color: color,
		value: value,
		class: class,
	}
}

func (c *badgePill) String() string {
	class := fmt.Sprintf("badge badge-pill badge-%s %s", c.color, c.class)

	return newSpan(class, c.value).String()
}
