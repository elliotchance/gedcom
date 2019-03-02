package html

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
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

func (c *IndividualInList) WriteHTMLTo(w io.Writer) (int64, error) {
	birthDate, birthPlace := c.individual.Birth()
	deathDate, deathPlace := c.individual.Death()

	birthPlaceName := prettyPlaceName(gedcom.String(birthPlace))
	deathPlaceName := prettyPlaceName(gedcom.String(deathPlace))

	birthDateText := core.NewText(gedcom.String(birthDate))
	deathDateText := core.NewText(gedcom.String(deathDate))

	link := NewIndividualLink(c.document, c.individual, c.visibility)
	birthPlaceLink := NewPlaceLink(c.document, birthPlaceName)
	deathPlaceLink := NewPlaceLink(c.document, deathPlaceName)
	birthLines := core.NewLines(birthDateText, birthPlaceLink)
	deathLines := core.NewLines(deathDateText, deathPlaceLink)

	return core.NewTableRow(
		core.NewTableCell(link).NoWrap(),
		core.NewTableCell(birthLines),
		core.NewTableCell(deathLines),
	).WriteHTMLTo(w)
}
