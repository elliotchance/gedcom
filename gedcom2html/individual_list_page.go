package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
)

// individualListPage is the page that lists of all the individuals.
type individualListPage struct {
	document *gedcom.Document
}

func newIndividualListPage(document *gedcom.Document) *individualListPage {
	return &individualListPage{
		document: document,
	}
}

func (c *individualListPage) String() string {
	table := []fmt.Stringer{
		newTableHead("Name", "Date of Birth", "Place of Birth", "Date of Death", "Place of Death"),
	}

	livingCount := 0
	for _, i := range c.document.Individuals() {
		if i.IsLiving() {
			livingCount += 1
			continue
		}

		table = append(table, newIndividualInList(c.document, i))
	}

	return newPage("People", newComponents(
		newHeader(c.document, "", selectedIndividualsTab),
		newRow(
			newColumn(entireRow, newText(fmt.Sprintf(
				"%d individuals are hidden because they are living.",
				livingCount,
			))),
		),
		newSpace(),
		newRow(
			newColumn(entireRow, newTable("", table...)),
		),
	)).String()
}
