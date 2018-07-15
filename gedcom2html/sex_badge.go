package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
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
	return fmt.Sprintf(`<span class="badge badge-%s">%s</span>`,
		colorClassForSex(c.sex), c.sex)
}
