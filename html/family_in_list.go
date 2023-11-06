package html

import (
	"io"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/gedcom/v39/html/core"
)

type FamilyInList struct {
	document   *gedcom.Document
	family     *gedcom.FamilyNode
	visibility LivingVisibility
	placesMap  map[string]*place
}

func NewFamilyInList(document *gedcom.Document, family *gedcom.FamilyNode, visibility LivingVisibility, placesMap map[string]*place) *FamilyInList {
	return &FamilyInList{
		document:   document,
		family:     family,
		visibility: visibility,
		placesMap:  placesMap,
	}
}

func (c *FamilyInList) WriteHTMLTo(w io.Writer) (int64, error) {
	date := "-"
	n := gedcom.First(gedcom.NodesWithTagPath(c.family, gedcom.TagMarriage, gedcom.TagDate))
	if n != nil {
		date = n.Value()
	}

	husband := NewIndividualLink(c.document, c.family.Husband().Individual(),
		c.visibility, c.placesMap)
	wife := NewIndividualLink(c.document, c.family.Wife().Individual(),
		c.visibility, c.placesMap)

	return core.NewTableRow(
		core.NewTableCell(husband),
		core.NewTableCell(core.NewText(date)).Class("text-center").NoWrap(),
		core.NewTableCell(wife),
	).WriteHTMLTo(w)
}
