package html

import (
	"github.com/elliotchance/gedcom"
	"io"
)

type IndividualStatistics struct {
	document   *gedcom.Document
	visibility LivingVisibility
}

func NewIndividualStatistics(document *gedcom.Document, visibility LivingVisibility) *IndividualStatistics {
	return &IndividualStatistics{
		document:   document,
		visibility: visibility,
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

	switch c.visibility {
	case LivingVisibilityShow, LivingVisibilityPlaceholder:
		// Proceed.

	case LivingVisibilityHide:
		// We need to pretend like there were never any living individuals in
		// the document at all.
		totalRow = NewKeyedTableRow("Total", NewNumber(total-living), true)
		livingRow = NewKeyedTableRow("Living", NewNumber(0), true)
		deadRow = NewKeyedTableRow("Dead", NewNumber(total-living), true)
	}

	s := NewComponents(totalRow, livingRow, deadRow)

	return NewCard(NewText("Individuals"), noBadgeCount, NewTable("", s)).WriteTo(w)
}

func keyedNumberRow(title string, total int) *KeyedTableRow {
	return NewKeyedTableRow(title, NewNumber(total), true)
}
