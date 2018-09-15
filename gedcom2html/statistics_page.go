package main

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
)

type statisticsPage struct {
	document          *gedcom.Document
	googleAnalyticsID string
}

func newStatisticsPage(document *gedcom.Document, googleAnalyticsID string) *statisticsPage {
	return &statisticsPage{
		document:          document,
		googleAnalyticsID: googleAnalyticsID,
	}
}

func (c *statisticsPage) String() string {
	return html.NewPage(
		"Statistics",
		html.NewComponents(
			newHeader(c.document, "", selectedStatisticsTab),
			html.NewBigTitle("Statistics"),
			html.NewSpace(),
			html.NewRow(
				html.NewColumn(html.HalfRow, html.NewComponents(
					newIndividualStatistics(c.document),
					html.NewSpace(),
					newFamilyStatistics(c.document),
					html.NewSpace(),
					newSourceStatistics(c.document),
					html.NewSpace(),
					newPlaceStatistics(c.document),
				)),
				html.NewColumn(html.HalfRow, newEventStatistics(c.document)),
			),
		),
		c.googleAnalyticsID,
	).String()
}
