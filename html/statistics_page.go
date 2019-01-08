package html

import (
	"github.com/elliotchance/gedcom"
	"io"
)

type StatisticsPage struct {
	document          *gedcom.Document
	googleAnalyticsID string
	options           PublishShowOptions
}

func NewStatisticsPage(document *gedcom.Document, googleAnalyticsID string, options PublishShowOptions) *StatisticsPage {
	return &StatisticsPage{
		document:          document,
		googleAnalyticsID: googleAnalyticsID,
		options:           options,
	}
}

func (c *StatisticsPage) WriteTo(w io.Writer) (int64, error) {
	return NewPage(
		"Statistics",
		NewComponents(
			NewPublishHeader(c.document, "", selectedStatisticsTab, c.options),
			NewBigTitle(1, NewText("Statistics")),
			NewSpace(),
			NewRow(
				NewColumn(HalfRow, NewComponents(
					NewIndividualStatistics(c.document),
					NewSpace(),
					NewFamilyStatistics(c.document),
					NewSpace(),
					NewSourceStatistics(c.document),
					NewSpace(),
					newPlaceStatistics(c.document),
				)),
				NewColumn(HalfRow, NewEventStatistics(c.document)),
			),
		),
		c.googleAnalyticsID,
	).WriteTo(w)
}
