package main

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
)

type placeStatistics struct {
	document *gedcom.Document
}

func newPlaceStatistics(document *gedcom.Document) *placeStatistics {
	return &placeStatistics{
		document: document,
	}
}

func (c *placeStatistics) String() string {
	total := html.NewNumber(len(getPlaces(c.document))).String()
	s := html.NewComponents(
		newKeyedTableRow("Total", total, true),
	)

	return newCard("Places", noBadgeCount, html.NewTable("", s)).String()
}
