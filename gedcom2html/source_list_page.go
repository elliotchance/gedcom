package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
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
		html.NewTableHead("Name"),
	}

	for _, source := range c.document.Sources() {
		table = append(table, newSourceInList(c.document, source))
	}

	return html.NewPage("Sources", html.NewComponents(
		newHeader(c.document, "", selectedSourcesTab),
		html.NewRow(
			html.NewColumn(html.EntireRow, html.NewTable("", table...)),
		),
	)).String()
}
