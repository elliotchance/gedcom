package main

import "github.com/elliotchance/gedcom/html"

// countBadge shows a pill badge containing an integer. The appropriate
// localization will be applied (like a thousands separator).
type countBadge struct {
	value int
}

func newCountBadge(value int) *countBadge {
	return &countBadge{
		value: value,
	}
}

func (c *countBadge) String() string {
	number := html.NewNumber(c.value)

	return newBadgePill("light", "", number.String()).String()
}
