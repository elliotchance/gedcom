package main

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
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

	return html.NewPage(
		name.String(),
		html.NewComponents(
			newHeader(c.document, name.String(), selectedExtraTab),
			newAllParentButtons(c.document, c.individual),
			html.NewBigTitle(html.NewIndividualName(c.individual, false, html.UnknownEmphasis).String()),
			html.NewHorizontalRuleRow(),
			html.NewRow(
				html.NewColumn(html.HalfRow, newIndividualNameAndSex(c.individual)),
				html.NewColumn(html.HalfRow, newIndividualAdditionalNames(c.individual)),
			),
			html.NewSpace(),
			newIndividualEvents(c.document, c.individual),
			html.NewSpace(),
			newPartnersAndChildren(c.document, c.individual),
		),
	).String()
}
