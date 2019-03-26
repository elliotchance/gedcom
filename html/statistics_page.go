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
}

func NewStatisticsPage(document *gedcom.Document, googleAnalyticsID string, options *PublishShowOptions, indexLetters []rune) *StatisticsPage {
	return &StatisticsPage{
		document:          document,
		googleAnalyticsID: googleAnalyticsID,
		options:           options,
		indexLetters:      indexLetters,
	}
}

func (c *StatisticsPage) WriteHTMLTo(w io.Writer) (int64, error) {
	return core.NewPage(
		"Statistics",
		core.NewComponents(
			NewPublishHeader(c.document, "", selectedStatisticsTab, c.options,
				c.indexLetters),
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
					newPlaceStatistics(c.document),
				)),
				core.NewColumn(core.HalfRow, NewEventStatistics(c.document)),
			),
		),
		c.googleAnalyticsID,
	).WriteHTMLTo(w)
}
