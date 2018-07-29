package main

import (
	"github.com/elliotchance/gedcom"
	"strconv"
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

	s := newComponents(
		newKeyedTableRow("Total", strconv.Itoa(total), true),
		newKeyedTableRow("Living", strconv.Itoa(living), true),
		newKeyedTableRow("Dead", strconv.Itoa(total-living), true),
	)

	return newCard("Individuals", noBadgeCount, newTable("", s)).String()
}
