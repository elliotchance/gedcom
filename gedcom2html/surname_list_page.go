package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
)

type surnameListPage struct {
	document          *gedcom.Document
	googleAnalyticsID string
}

func newSurnameListPage(document *gedcom.Document, googleAnalyticsID string) *surnameListPage {
	return &surnameListPage{
		document:          document,
		googleAnalyticsID: googleAnalyticsID,
	}
}

func (c *surnameListPage) String() string {
	table := []fmt.Stringer{
		html.NewTableHead("Surname", "Number of Individuals"),
	}

	for _, surname := range getSurnames(c.document) {
		table = append(table, newSurnameInList(c.document, surname))
	}

	return html.NewPage("Surnames", html.NewComponents(
		newHeader(c.document, "", selectedSurnamesTab),
		html.NewRow(
			html.NewColumn(html.EntireRow, html.NewTable("", table...)),
		),
	), c.googleAnalyticsID).String()
}
