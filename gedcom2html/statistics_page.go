package main

import (
	"github.com/elliotchance/gedcom"
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
	return newPage(
		"Statistics",
		newComponents(
			newHeader(c.document, "", selectedStatisticsTab),
			newBigName("Statistics"),
			newSpace(),
			newRow(
				newColumn(halfRow, newIndividualStatistics(c.document)),
				newColumn(halfRow, newEventStatistics(c.document)),
			),
			newSpace(),
			newRow(
				newColumn(halfRow, newFamilyStatistics(c.document)),
				newColumn(halfRow, newSourceStatistics(c.document)),
			),
			newSpace(),
			newRow(
				newColumn(halfRow, newPlaceStatistics(c.document)),
				newColumn(halfRow, newEmpty()),
			),
		),
	).String()
}
