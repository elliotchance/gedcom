package main

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

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
	p := message.NewPrinter(language.English)

	return newBadgePill("light", "", p.Sprintf("%d", c.value)).String()
}
