package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
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
		html.NewTableHead("Husband", "Date", "Wife"),
	}

	for _, family := range c.document.Families() {
		table = append(table, newFamilyInList(c.document, family))
	}

	return html.NewPage("Families", html.NewComponents(
		newHeader(c.document, "", selectedFamiliesTab),
		html.NewRow(
			html.NewColumn(html.EntireRow, html.NewTable("", table...)),
		),
	)).String()
}
