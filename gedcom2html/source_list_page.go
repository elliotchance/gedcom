package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
)

type sourceListPage struct {
	document *gedcom.Document
}

func newSourceListPage(document *gedcom.Document) *sourceListPage {
	return &sourceListPage{
		document: document,
	}
}

func (c *sourceListPage) String() string {
	table := []fmt.Stringer{
		newTableHead("Name"),
	}

	for _, source := range c.document.Sources() {
		table = append(table, newSourceInList(c.document, source))
	}

	return newPage("Sources", newComponents(
		newHeader(c.document, "", selectedSourcesTab),
		newRow(
			newColumn(entireRow, newTable("", table...)),
		),
	)).String()
}
