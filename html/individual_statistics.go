package html

import (
	"io"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/gedcom/v39/html/core"
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

func (c *IndividualStatistics) WriteHTMLTo(w io.Writer) (int64, error) {
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
		totalRow = core.NewKeyedTableRow("Total", core.NewNumber(total-living), true)
		livingRow = core.NewKeyedTableRow("Living", core.NewNumber(0), true)
		deadRow = core.NewKeyedTableRow("Dead", core.NewNumber(total-living), true)
	}

	s := core.NewComponents(totalRow, livingRow, deadRow)

	return core.NewCard(core.NewText("Individuals"), core.CardNoBadgeCount,
		core.NewTable("", s)).WriteHTMLTo(w)
}

func keyedNumberRow(title string, total int) *core.KeyedTableRow {
	return core.NewKeyedTableRow(title, core.NewNumber(total), true)
}
