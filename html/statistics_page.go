package html

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
	"io"
)

type StatisticsPage struct {
	document          *gedcom.Document
	googleAnalyticsID string
	options           *PublishShowOptions
	indexLetters      []rune
	placesMap         map[string]*place
}

func NewStatisticsPage(document *gedcom.Document, googleAnalyticsID string, options *PublishShowOptions, indexLetters []rune, placesMap map[string]*place) *StatisticsPage {
	return &StatisticsPage{
		document:          document,
		googleAnalyticsID: googleAnalyticsID,
		options:           options,
		indexLetters:      indexLetters,
		placesMap:         placesMap,
	}
}

func (c *StatisticsPage) WriteHTMLTo(w io.Writer) (int64, error) {
	return core.NewPage(
		"Statistics",
		core.NewComponents(
			NewPublishHeader(c.document, "", selectedStatisticsTab, c.options,
				c.indexLetters, c.placesMap),
			core.NewBigTitle(1, core.NewText("Statistics")),
			core.NewSpace(),
			core.NewRow(
				core.NewColumn(core.HalfRow, core.NewComponents(
					NewIndividualStatistics(c.document, c.options.LivingVisibility),
					core.NewSpace(),
					NewFamilyStatistics(c.document),
					core.NewSpace(),
					NewSourceStatistics(c.document),
					core.NewSpace(),
					newPlaceStatistics(c.document, c.placesMap),
				)),
				core.NewColumn(core.HalfRow, NewEventStatistics(c.document)),
			),
		),
		c.googleAnalyticsID,
	).WriteHTMLTo(w)
}
