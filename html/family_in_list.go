package html

import (
	"github.com/elliotchance/gedcom"
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

func (c *FamilyInList) WriteTo(w io.Writer) (int64, error) {
	date := "-"
	n := gedcom.First(gedcom.NodesWithTagPath(c.family, gedcom.TagMarriage, gedcom.TagDate))
	if n != nil {
		date = n.Value()
	}

	husband := NewIndividualLink(c.document, c.family.Husband().Individual(), c.visibility)
	wife := NewIndividualLink(c.document, c.family.Wife().Individual(), c.visibility)

	return NewTableRow(
		NewTableCell(husband),
		NewTableCell(NewText(date)).Class("text-center").NoWrap(),
		NewTableCell(wife),
	).WriteTo(w)
}
