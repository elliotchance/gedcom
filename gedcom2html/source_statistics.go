package main

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
	"strconv"
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
	s := html.NewComponents(
		newKeyedTableRow("Total", strconv.Itoa(len(c.document.Sources())), true),
	)

	return newCard("Sources", noBadgeCount, html.NewTable("", s)).String()
}
