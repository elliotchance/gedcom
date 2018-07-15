package main

import "fmt"

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
	return fmt.Sprintf(`<span class="badge badge-pill badge-%s %s">%s</span>`,
		c.color, c.class, c.value)
}
