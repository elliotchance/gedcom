package html

import (
	"io"
	"sort"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/gedcom/v39/html/core"
)

type EventStatistics struct {
	document *gedcom.Document
}

func NewEventStatistics(document *gedcom.Document) *EventStatistics {
	return &EventStatistics{
		document: document,
	}
}

func (c *EventStatistics) WriteHTMLTo(w io.Writer) (int64, error) {
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

	rows := []core.Component{
		core.NewKeyedTableRow("Total", core.NewNumber(total), true),
	}

	keys := []string{}
	for name := range counts {
		keys = append(keys, name)
	}

	sort.Strings(keys)

	for _, name := range keys {
		number := core.NewNumber(counts[name])
		tableRow := core.NewKeyedTableRow(name, number, true)
		rows = append(rows, tableRow)
	}

	table := core.NewTable("", core.NewComponents(rows...))

	return core.NewCard(core.NewText("Events"), core.CardNoBadgeCount, table).
		WriteHTMLTo(w)
}
