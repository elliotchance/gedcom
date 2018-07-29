package main

import (
	"github.com/elliotchance/gedcom"
	"strconv"
)

type eventStatistics struct {
	document *gedcom.Document
}

func newEventStatistics(document *gedcom.Document) *eventStatistics {
	return &eventStatistics{
		document: document,
	}
}

func (c *eventStatistics) String() string {
	births := 0
	christenings := 0
	deaths := 0
	burials := 0

	for _, individual := range c.document.Individuals() {
		if date, _ := getBirth(individual); date != "" {
			births += 1
		}

		if date, _ := getDeath(individual); date != "" {
			deaths += 1
		}

		if n := individual.FirstNodeWithTagPath(gedcom.TagChristening); n != nil {
			christenings += 1
		}

		if n := individual.FirstNodeWithTagPath(gedcom.TagBurial); n != nil {
			burials += 1
		}
	}

	s := newComponents(
		newKeyedTableRow("Total", strconv.Itoa(births+christenings+deaths+burials), true),
		newKeyedTableRow("Births", strconv.Itoa(births), true),
		newKeyedTableRow("Christenings", strconv.Itoa(christenings), true),
		newKeyedTableRow("Deaths", strconv.Itoa(deaths), true),
		newKeyedTableRow("Burials", strconv.Itoa(burials), true),
	)

	return newCard("Events", noBadgeCount, newTable("", s)).String()
}
