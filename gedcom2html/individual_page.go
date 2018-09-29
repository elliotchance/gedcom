package main

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
)

// individualPage is the page that shows detailed information about an
// individual.
type individualPage struct {
	document          *gedcom.Document
	individual        *gedcom.IndividualNode
	googleAnalyticsID string
}

func newIndividualPage(document *gedcom.Document, individual *gedcom.IndividualNode, googleAnalyticsID string) *individualPage {
	return &individualPage{
		document:          document,
		individual:        individual,
		googleAnalyticsID: googleAnalyticsID,
	}
}

func (c *individualPage) String() string {
	name := c.individual.Names()[0]

	individualName := html.NewIndividualName(c.individual, false,
		html.UnknownEmphasis)
	individualDates := html.NewIndividualDates(c.individual, false)

	return html.NewPage(
		name.String(),
		html.NewComponents(
			newHeader(c.document, name.String(), selectedExtraTab),
			newAllParentButtons(c.document, c.individual),
			html.NewBigTitle(1, individualName.String()),
			html.NewBigTitle(3, individualDates.String()),
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
		c.googleAnalyticsID,
	).String()
}
