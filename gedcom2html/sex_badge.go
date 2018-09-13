package main

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
)

// sexBadge shows a coloured "Male", "Female" or "Unknown" badge.
type sexBadge struct {
	sex gedcom.Sex
}

func newSexBadge(sex gedcom.Sex) *sexBadge {
	return &sexBadge{
		sex: sex,
	}
}

func (c *sexBadge) String() string {
	return html.Sprintf(`<span class="badge badge-%s">%s</span>`,
		colorClassForSex(c.sex), c.sex)
}
