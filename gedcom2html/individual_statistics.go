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

	totalRow := newKeyedTableRow("Total", html.NewNumber(total).String(), true)
	livingRow := newKeyedTableRow("Living", html.NewNumber(living).String(), true)
	deadRow := newKeyedTableRow("Dead", html.NewNumber(total-living).String(), true)

	s := html.NewComponents(totalRow, livingRow, deadRow)

	return newCard("Individuals", noBadgeCount, html.NewTable("", s)).String()
}
