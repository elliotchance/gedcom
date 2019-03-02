package core

import "io"

// CountBadge shows a pill badge containing an integer. The appropriate
// localization will be applied (like a thousands separator).
type CountBadge struct {
	value int
}

func NewCountBadge(value int) *CountBadge {
	return &CountBadge{
		value: value,
	}
}

func (c *CountBadge) WriteHTMLTo(w io.Writer) (int64, error) {
	number := NewNumber(c.value)

	return NewBadgePill("light", "", number).WriteHTMLTo(w)
}
