package main

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
)

type familyInList struct {
	document *gedcom.Document
	family   *gedcom.FamilyNode
}

func newFamilyInList(document *gedcom.Document, family *gedcom.FamilyNode) *familyInList {
	return &familyInList{
		document: document,
		family:   family,
	}
}

func (c *familyInList) String() string {
	date := "-"
	n := gedcom.First(gedcom.NodesWithTagPath(c.family, gedcom.TagMarriage, gedcom.TagDate))
	if n != nil {
		date = n.Value()
	}

	husband := newIndividualLink(c.document, c.family.Husband())
	wife := newIndividualLink(c.document, c.family.Wife())

	return html.NewTableRow(
		html.NewTableCell(husband),
		html.NewTableCell(html.NewText(date)).Class("text-center").NoWrap(),
		html.NewTableCell(wife),
	).String()
}
