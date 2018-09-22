package main

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
)

type sourceStatistics struct {
	document *gedcom.Document
}

func newSourceStatistics(document *gedcom.Document) *sourceStatistics {
	return &sourceStatistics{
		document: document,
	}
}

func (c *sourceStatistics) String() string {
	total := html.NewNumber(len(c.document.Sources())).String()
	s := html.NewComponents(
		newKeyedTableRow("Total", total, true),
	)

	return newCard("Sources", noBadgeCount, html.NewTable("", s)).String()
}
