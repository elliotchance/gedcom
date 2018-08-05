package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"sort"
	"strings"
)

// individualListPage is the page that lists of all the individuals.
type individualListPage struct {
	document       *gedcom.Document
	selectedLetter rune
}

func newIndividualListPage(document *gedcom.Document, selectedLetter rune) *individualListPage {
	return &individualListPage{
		document:       document,
		selectedLetter: selectedLetter,
	}
}

func (c *individualListPage) String() string {
	table := []fmt.Stringer{
		newTableHead("Name", "Birth", "Death"),
	}

	individuals := gedcom.IndividualNodes{}

	for _, individual := range c.document.Individuals() {
		name := strings.ToLower(individual.Name().String())
		if name == "" {
			name = "#"
		}

		if rune(name[0]) == c.selectedLetter {
			individuals = append(individuals, individual)
		}
	}

	// Sort individuals by name.
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
		newIndividualIndexHeader(c.document, c.selectedLetter),
		newSpace(),
		newRow(
			newColumn(entireRow, newTable("", table...)),
		),
	)).String()
}
