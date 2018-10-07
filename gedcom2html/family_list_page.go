package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
)

type familyListPage struct {
	document          *gedcom.Document
	googleAnalyticsID string
}

func newFamilyListPage(document *gedcom.Document, googleAnalyticsID string) *familyListPage {
	return &familyListPage{
		document:          document,
		googleAnalyticsID: googleAnalyticsID,
	}
}

func (c *familyListPage) String() string {
	table := []fmt.Stringer{
		html.NewTableHead("Husband", "Date", "Wife"),
	}

	for _, family := range c.document.Families() {
		familyInList := newFamilyInList(c.document, family)
		table = append(table, familyInList)
	}

	column := html.NewColumn(html.EntireRow, html.NewTable("", table...))
	header := newHeader(c.document, "", selectedFamiliesTab)
	components := html.NewComponents(header, html.NewRow(column))

	return html.NewPage("Families", components, c.googleAnalyticsID).String()
}
