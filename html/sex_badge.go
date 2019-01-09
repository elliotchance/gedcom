package html

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"io"
)

// SexBadge shows a coloured "Male", "Female" or "Unknown" badge.
type SexBadge struct {
	sex gedcom.Sex
}

func NewSexBadge(sex gedcom.Sex) *SexBadge {
	return &SexBadge{
		sex: sex,
	}
}

func (c *SexBadge) WriteTo(w io.Writer) (int64, error) {
	return NewSpan(
		fmt.Sprintf("badge badge-%s", colorClassForSex(c.sex)),
		NewText(c.sex.String()),
	).WriteTo(w)
}
