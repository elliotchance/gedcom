package main

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
)

// individualInList is a single row in the table of individuals on the list
// page.
type individualInList struct {
	individual *gedcom.IndividualNode
	document   *gedcom.Document
}

func newIndividualInList(document *gedcom.Document, individual *gedcom.IndividualNode) *individualInList {
	return &individualInList{
		individual: individual,
		document:   document,
	}
}

func (c *individualInList) String() string {
	birthDate, birthPlace := c.individual.Birth()
	deathDate, deathPlace := c.individual.Death()

	birthPlaceName := prettyPlaceName(birthPlace.String())
	deathPlaceName := prettyPlaceName(deathPlace.String())

	birthDateText := html.NewText(birthDate.String())
	deathDateText := html.NewText(deathDate.String())

	link := newIndividualLink(c.document, c.individual)
	birthPlaceLink := newPlaceLink(c.document, birthPlaceName)
	deathPlaceLink := newPlaceLink(c.document, deathPlaceName)
	birthLines := html.NewLines(birthDateText, birthPlaceLink)
	deathLines := html.NewLines(deathDateText, deathPlaceLink)

	return html.NewTableRow(
		html.NewTableCell(link).NoWrap(),
		html.NewTableCell(birthLines),
		html.NewTableCell(deathLines),
	).String()
}
