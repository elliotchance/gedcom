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
	visibility        LivingVisibility
}

func NewIndividualPage(document *gedcom.Document, individual *gedcom.IndividualNode, googleAnalyticsID string, options PublishShowOptions, visibility LivingVisibility) *IndividualPage {
	return &IndividualPage{
		document:          document,
		individual:        individual,
		googleAnalyticsID: googleAnalyticsID,
		options:           options,
		visibility:        visibility,
	}
}

func (c *IndividualPage) WriteTo(w io.Writer) (int64, error) {
	name := c.individual.Names()[0]

	individualName := NewIndividualName(c.individual, c.visibility,
		UnknownEmphasis)
	individualDates := NewIndividualDates(c.individual, c.visibility)

	return NewPage(
		name.String(),
		NewComponents(
			NewPublishHeader(c.document, name.String(), selectedExtraTab, c.options),
			NewAllParentButtons(c.document, c.individual, c.visibility),
			NewBigTitle(1, individualName),
			NewBigTitle(3, individualDates),
			NewHorizontalRuleRow(),
			NewRow(
				NewColumn(HalfRow, NewIndividualNameAndSex(c.individual)),
				NewColumn(HalfRow, NewIndividualAdditionalNames(c.individual)),
			),
			NewSpace(),
			newIndividualEvents(c.document, c.individual, c.visibility),
			NewSpace(),
			NewPartnersAndChildren(c.document, c.individual, c.visibility),
		),
		c.googleAnalyticsID,
	).WriteTo(w)
}
