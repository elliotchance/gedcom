package main

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
)

type statisticsPage struct {
	document *gedcom.Document
}

func newStatisticsPage(document *gedcom.Document) *statisticsPage {
	return &statisticsPage{
		document: document,
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
				html.NewColumn(html.HalfRow, newIndividualStatistics(c.document)),
				html.NewColumn(html.HalfRow, newEventStatistics(c.document)),
			),
			html.NewSpace(),
			html.NewRow(
				html.NewColumn(html.HalfRow, newFamilyStatistics(c.document)),
				html.NewColumn(html.HalfRow, newSourceStatistics(c.document)),
			),
			html.NewSpace(),
			html.NewRow(
				html.NewColumn(html.HalfRow, newPlaceStatistics(c.document)),
				html.NewColumn(html.HalfRow, newEmpty()),
			),
		),
	).String()
}
