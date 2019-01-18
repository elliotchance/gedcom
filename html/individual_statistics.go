package html

import (
	"github.com/elliotchance/gedcom"
	"io"
)

type IndividualStatistics struct {
	document *gedcom.Document
}

func NewIndividualStatistics(document *gedcom.Document) *IndividualStatistics {
	return &IndividualStatistics{
		document: document,
	}
}

func (c *IndividualStatistics) WriteTo(w io.Writer) (int64, error) {
	total := 0
	living := 0

	for _, individual := range c.document.Individuals() {
		total += 1

		if individual.IsLiving() {
			living += 1
		}
	}

	totalRow := keyedNumberRow("Total", total)
	livingRow := keyedNumberRow("Living", living)
	deadRow := keyedNumberRow("Dead", total-living)

	s := NewComponents(totalRow, livingRow, deadRow)

	return NewCard("Individuals", noBadgeCount, NewTable("", s)).WriteTo(w)
}

func keyedNumberRow(title string, total int) *KeyedTableRow {
	return NewKeyedTableRow(title, NewNumber(total), true)
}
