package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
	"sort"
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
	counts := map[string]int{}

	for _, individual := range c.document.Individuals() {
		for _, event := range individual.AllEvents() {
			counts[event.Tag().String()] += 1
		}
	}

	total := 0
	for _, count := range counts {
		total += count
	}

	rows := []fmt.Stringer{
		newKeyedTableRow("Total", strconv.Itoa(total), true),
	}

	keys := []string{}
	for name := range counts {
		keys = append(keys, name)
	}

	sort.Strings(keys)

	for _, name := range keys {
		rows = append(rows,
			newKeyedTableRow(name, strconv.Itoa(counts[name]), true))
	}

	return newCard("Events", noBadgeCount,
		html.NewTable("", html.NewComponents(rows...))).String()
}
