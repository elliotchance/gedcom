package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
)

type familyListPage struct {
	document *gedcom.Document
}

func newFamilyListPage(document *gedcom.Document) *familyListPage {
	return &familyListPage{
		document: document,
	}
}

func (c *familyListPage) String() string {
	table := []fmt.Stringer{
		newTableHead("Husband", "Date", "Wife"),
	}

	for _, family := range c.document.Families() {
		table = append(table, newFamilyInList(c.document, family))
	}

	return newPage("Families", newComponents(
		newHeader(c.document, "", selectedFamiliesTab),
		newRow(
			newColumn(entireRow, newTable("", table...)),
		),
	)).String()
}
