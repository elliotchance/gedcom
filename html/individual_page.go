package html

import (
	"github.com/elliotchance/gedcom"
	"io"
)

// IndividualPage is the page that shows detailed information about an
// individual.
type IndividualPage struct {
	document          *gedcom.Document
	individual        *gedcom.IndividualNode
	googleAnalyticsID string
	options           PublishShowOptions
}

func NewIndividualPage(document *gedcom.Document, individual *gedcom.IndividualNode, googleAnalyticsID string, options PublishShowOptions) *IndividualPage {
	return &IndividualPage{
		document:          document,
		individual:        individual,
		googleAnalyticsID: googleAnalyticsID,
		options:           options,
	}
}

func (c *IndividualPage) WriteTo(w io.Writer) (int64, error) {
	name := c.individual.Names()[0]

	individualName := NewIndividualName(c.individual, false,
		UnknownEmphasis)
	individualDates := NewIndividualDates(c.individual, false)

	return NewPage(
		name.String(),
		NewComponents(
			NewPublishHeader(c.document, name.String(), selectedExtraTab, c.options),
			NewAllParentButtons(c.document, c.individual),
			NewBigTitle(1, individualName),
			NewBigTitle(3, individualDates),
			NewHorizontalRuleRow(),
			NewRow(
				NewColumn(HalfRow, NewIndividualNameAndSex(c.individual)),
				NewColumn(HalfRow, NewIndividualAdditionalNames(c.individual)),
			),
			NewSpace(),
			newIndividualEvents(c.document, c.individual),
			NewSpace(),
			NewPartnersAndChildren(c.document, c.individual),
		),
		c.googleAnalyticsID,
	).WriteTo(w)
}
