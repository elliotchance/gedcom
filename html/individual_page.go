package html

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
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

func (c *IndividualPage) WriteHTMLTo(w io.Writer) (int64, error) {
	name := c.individual.Names()[0]

	individualName := NewIndividualName(c.individual, c.visibility,
		UnknownEmphasis)
	individualDates := NewIndividualDates(c.individual, c.visibility)

	return core.NewPage(
		name.String(),
		core.NewComponents(
			NewPublishHeader(c.document, name.String(), selectedExtraTab, c.options),
			NewAllParentButtons(c.document, c.individual, c.visibility),
			core.NewBigTitle(1, individualName),
			core.NewBigTitle(3, individualDates),
			core.NewHorizontalRuleRow(),
			core.NewRow(
				core.NewColumn(core.HalfRow, NewIndividualNameAndSex(c.individual)),
				core.NewColumn(core.HalfRow, NewIndividualAdditionalNames(c.individual)),
			),
			core.NewSpace(),
			newIndividualEvents(c.document, c.individual, c.visibility),
			core.NewSpace(),
			NewPartnersAndChildren(c.document, c.individual, c.visibility),
		),
		c.googleAnalyticsID,
	).WriteHTMLTo(w)
}
