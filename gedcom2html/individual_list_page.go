package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"sort"
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
		newTableHead("Name", "Birth", "Death"),
	}

	// Sort individuals by name.
	individuals := c.document.Individuals()
	sort.Slice(individuals, func(i, j int) bool {
		return individuals[i].Name().String() < individuals[j].Name().String()
	})

	livingCount := 0
	for _, i := range individuals {
		if i.IsLiving() {
			livingCount += 1
			continue
		}

		table = append(table, newIndividualInList(c.document, i))
	}

	return newPage("Individuals", newComponents(
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
