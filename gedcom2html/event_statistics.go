package main

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
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
		if date, _ := html.GetBirth(individual); date != "" {
			births += 1
		}

		if date, _ := html.GetDeath(individual); date != "" {
			deaths += 1
		}

		if n := gedcom.First(gedcom.NodesWithTagPath(individual, gedcom.TagChristening)); n != nil {
			christenings += 1
		}

		if n := gedcom.First(gedcom.NodesWithTagPath(individual, gedcom.TagBurial)); n != nil {
			burials += 1
		}
	}

	s := html.NewComponents(
		newKeyedTableRow("Total", strconv.Itoa(births+christenings+deaths+burials), true),
		newKeyedTableRow("Births", strconv.Itoa(births), true),
		newKeyedTableRow("Christenings", strconv.Itoa(christenings), true),
		newKeyedTableRow("Deaths", strconv.Itoa(deaths), true),
		newKeyedTableRow("Burials", strconv.Itoa(burials), true),
	)

	return newCard("Events", noBadgeCount, html.NewTable("", s)).String()
}
