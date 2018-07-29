package main

import (
	"github.com/elliotchance/gedcom"
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
	s := newComponents(
		newKeyedTableRow("Total", strconv.Itoa(len(c.document.Sources())), true),
	)

	return newCard("Sources", noBadgeCount, newTable("", s)).String()
}
