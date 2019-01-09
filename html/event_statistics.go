package html

import (
	"github.com/elliotchance/gedcom"
	"io"
	"sort"
)

type EventStatistics struct {
	document *gedcom.Document
}

func NewEventStatistics(document *gedcom.Document) *EventStatistics {
	return &EventStatistics{
		document: document,
	}
}

func (c *EventStatistics) WriteTo(w io.Writer) (int64, error) {
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

	rows := []Component{
		NewKeyedTableRow("Total", NewNumber(total), true),
	}

	keys := []string{}
	for name := range counts {
		keys = append(keys, name)
	}

	sort.Strings(keys)

	for _, name := range keys {
		number := NewNumber(counts[name])
		tableRow := NewKeyedTableRow(name, number, true)
		rows = append(rows, tableRow)
	}

	table := NewTable("", NewComponents(rows...))

	return NewCard("Events", noBadgeCount, table).WriteTo(w)
}
