package html

import (
	"fmt"
	"io"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/gedcom/v39/html/core"
)

// SexBadge shows a coloured "Male", "Female" or "Unknown" badge.
type SexBadge struct {
	sex *gedcom.SexNode
}

func NewSexBadge(sex *gedcom.SexNode) *SexBadge {
	return &SexBadge{
		sex: sex,
	}
}

func (c *SexBadge) WriteHTMLTo(w io.Writer) (int64, error) {
	return core.NewSpan(
		fmt.Sprintf("badge badge-%s", colorClassForSex(c.sex)),
		core.NewText(c.sex.String()),
	).WriteHTMLTo(w)
}
