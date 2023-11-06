package html

import (
	"io"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/gedcom/v39/html/core"
)

// IndividualPage is the page that shows detailed information about an
// individual.
type IndividualPage struct {
	document          *gedcom.Document
	individual        *gedcom.IndividualNode
	googleAnalyticsID string
	options           *PublishShowOptions
	indexLetters      []rune
	placesMap         map[string]*place
}

func NewIndividualPage(document *gedcom.Document, individual *gedcom.IndividualNode, googleAnalyticsID string, options *PublishShowOptions, indexLetters []rune, placesMap map[string]*place) *IndividualPage {
	return &IndividualPage{
		document:          document,
		individual:        individual,
		googleAnalyticsID: googleAnalyticsID,
		options:           options,
		indexLetters:      indexLetters,
		placesMap:         placesMap,
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
			NewPublishHeader(c.document, name.String(), selectedExtraTab,
				c.options, c.indexLetters, c.placesMap),
			NewAllParentButtons(c.document, c.individual,
				c.options.LivingVisibility, c.placesMap),
			core.NewBigTitle(1, individualName),
			core.NewBigTitle(3, individualDates),
			core.NewHorizontalRuleRow(),
			core.NewRow(
				core.NewColumn(core.HalfRow, NewIndividualNameAndSex(c.individual)),
				core.NewColumn(core.HalfRow, NewIndividualAdditionalNames(c.individual)),
			),
			core.NewSpace(),
			NewIndividualEvents(c.document, c.individual,
				c.options.LivingVisibility, c.placesMap),
			core.NewSpace(),
			NewPartnersAndChildren(c.document, c.individual,
				c.options.LivingVisibility, c.placesMap),
		),
		c.googleAnalyticsID,
	).WriteHTMLTo(w)
}
