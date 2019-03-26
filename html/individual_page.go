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
	options           *PublishShowOptions
	indexLetters      []rune
}

func NewIndividualPage(document *gedcom.Document, individual *gedcom.IndividualNode, googleAnalyticsID string, options *PublishShowOptions, indexLetters []rune) *IndividualPage {
	return &IndividualPage{
		document:          document,
		individual:        individual,
		googleAnalyticsID: googleAnalyticsID,
		options:           options,
		indexLetters:      indexLetters,
	}
}

func (c *IndividualPage) WriteHTMLTo(w io.Writer) (int64, error) {
	name := c.individual.Names()[0]

	individualName := NewIndividualName(c.individual, c.options.LivingVisibility,
		UnknownEmphasis)
	individualDates := NewIndividualDates(c.individual, c.options.LivingVisibility)

	return core.NewPage(
		name.String(),
		core.NewComponents(
			NewPublishHeader(c.document, name.String(), selectedExtraTab, c.options, c.indexLetters),
			NewAllParentButtons(c.document, c.individual, c.options.LivingVisibility),
			core.NewBigTitle(1, individualName),
			core.NewBigTitle(3, individualDates),
			core.NewHorizontalRuleRow(),
			core.NewRow(
				core.NewColumn(core.HalfRow, NewIndividualNameAndSex(c.individual)),
				core.NewColumn(core.HalfRow, NewIndividualAdditionalNames(c.individual)),
			),
			core.NewSpace(),
			newIndividualEvents(c.document, c.individual, c.options.LivingVisibility),
			core.NewSpace(),
			NewPartnersAndChildren(c.document, c.individual, c.options.LivingVisibility),
		),
		c.googleAnalyticsID,
	).WriteHTMLTo(w)
}
