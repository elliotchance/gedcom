package html

import (
	"io"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/gedcom/v39/html/core"
)

// IndividualInList is a single row in the table of individuals on the list
// page.
type IndividualInList struct {
	individual *gedcom.IndividualNode
	document   *gedcom.Document
	visibility LivingVisibility
	placesMap  map[string]*place
}

func NewIndividualInList(document *gedcom.Document, individual *gedcom.IndividualNode, visibility LivingVisibility, placesMap map[string]*place) *IndividualInList {
	return &IndividualInList{
		individual: individual,
		document:   document,
		visibility: visibility,
		placesMap:  placesMap,
	}
}

func (c *IndividualInList) WriteHTMLTo(w io.Writer) (int64, error) {
	birthDate, birthPlace := c.individual.Birth()
	deathDate, deathPlace := c.individual.Death()

	birthPlaceName := prettyPlaceName(gedcom.String(birthPlace))
	deathPlaceName := prettyPlaceName(gedcom.String(deathPlace))

	birthDateText := core.NewText(gedcom.String(birthDate))
	deathDateText := core.NewText(gedcom.String(deathDate))

	link := NewIndividualLink(c.document, c.individual, c.visibility, c.placesMap)
	birthPlaceLink := NewPlaceLink(c.document, birthPlaceName, c.placesMap)
	deathPlaceLink := NewPlaceLink(c.document, deathPlaceName, c.placesMap)
	birthLines := core.NewLines(birthDateText, birthPlaceLink)
	deathLines := core.NewLines(deathDateText, deathPlaceLink)

	return core.NewTableRow(
		core.NewTableCell(link).NoWrap(),
		core.NewTableCell(birthLines),
		core.NewTableCell(deathLines),
	).WriteHTMLTo(w)
}
