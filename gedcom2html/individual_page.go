package main

import (
	"github.com/elliotchance/gedcom"
)

// individualPage is the page that shows detailed information about an
// individual.
type individualPage struct {
	document   *gedcom.Document
	individual *gedcom.IndividualNode
}

func newIndividualPage(document *gedcom.Document, individual *gedcom.IndividualNode) *individualPage {
	return &individualPage{
		document:   document,
		individual: individual,
	}
}

func (c *individualPage) String() string {
	name := c.individual.Names()[0]

	return newPage(
		name.String(),
		newComponents(
			newHeader(c.document, name.String(), selectedExtraTab),
			newAllParentButtons(c.document, c.individual),
			newBigIndividualName(c.individual),
			newHorizontalRuleRow(),
			newRow(
				newColumn(halfRow, newIndividualNameAndSex(c.individual)),
				newColumn(halfRow, newIndividualAdditionalNames(c.individual)),
			),
			newSpace(),
			newIndividualEvents(c.document, c.individual),
			newSpace(),
			newPartnersAndChildren(c.document, c.individual),
		),
	).String()
}
