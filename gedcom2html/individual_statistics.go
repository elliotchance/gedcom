package main

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
)

type individualStatistics struct {
	document *gedcom.Document
}

func newIndividualStatistics(document *gedcom.Document) *individualStatistics {
	return &individualStatistics{
		document: document,
	}
}

func (c *individualStatistics) String() string {
	total := 0
	living := 0

	for _, individual := range c.document.Individuals() {
		total += 1

		if individual.IsLiving() {
			living += 1
		}
	}

	s := html.NewComponents(
		newKeyedTableRow("Total", html.NewNumber(total).String(), true),
		newKeyedTableRow("Living", html.NewNumber(living).String(), true),
		newKeyedTableRow("Dead", html.NewNumber(total-living).String(), true),
	)

	return newCard("Individuals", noBadgeCount, html.NewTable("", s)).String()
}
