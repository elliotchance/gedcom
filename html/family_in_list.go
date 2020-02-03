package html

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
	"github.com/elliotchance/gedcom/tag"
	"io"
)

type FamilyInList struct {
	document   *gedcom.Document
	family     *gedcom.FamilyNode
	visibility LivingVisibility
}

func NewFamilyInList(document *gedcom.Document, family *gedcom.FamilyNode, visibility LivingVisibility) *FamilyInList {
	return &FamilyInList{
		document:   document,
		family:     family,
		visibility: visibility,
	}
}

func (c *FamilyInList) WriteHTMLTo(w io.Writer) (int64, error) {
	date := "-"
	n := gedcom.First(gedcom.NodesWithTagPath(c.family, tag.TagMarriage, tag.TagDate))
	if n != nil {
		date = n.Value()
	}

	husband := NewIndividualLink(c.document, c.family.Husband().Individual(), c.visibility)
	wife := NewIndividualLink(c.document, c.family.Wife().Individual(), c.visibility)

	return core.NewTableRow(
		core.NewTableCell(husband),
		core.NewTableCell(core.NewText(date)).Class("text-center").NoWrap(),
		core.NewTableCell(wife),
	).WriteHTMLTo(w)
}
