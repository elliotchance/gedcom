package html

import (
	"github.com/elliotchance/gedcom"
	"io"
)

// IndividualInList is a single row in the table of individuals on the list
// page.
type IndividualInList struct {
	individual *gedcom.IndividualNode
	document   *gedcom.Document
	visibility LivingVisibility
}

func NewIndividualInList(document *gedcom.Document, individual *gedcom.IndividualNode, visibility LivingVisibility) *IndividualInList {
	return &IndividualInList{
		individual: individual,
		document:   document,
		visibility: visibility,
	}
}

func (c *IndividualInList) WriteTo(w io.Writer) (int64, error) {
	birthDate, birthPlace := c.individual.Birth()
	deathDate, deathPlace := c.individual.Death()

	birthPlaceName := prettyPlaceName(gedcom.String(birthPlace))
	deathPlaceName := prettyPlaceName(gedcom.String(deathPlace))

	birthDateText := NewText(gedcom.String(birthDate))
	deathDateText := NewText(gedcom.String(deathDate))

	link := NewIndividualLink(c.document, c.individual, c.visibility)
	birthPlaceLink := NewPlaceLink(c.document, birthPlaceName)
	deathPlaceLink := NewPlaceLink(c.document, deathPlaceName)
	birthLines := NewLines(birthDateText, birthPlaceLink)
	deathLines := NewLines(deathDateText, deathPlaceLink)

	return NewTableRow(
		NewTableCell(link).NoWrap(),
		NewTableCell(birthLines),
		NewTableCell(deathLines),
	).WriteTo(w)
}
