package main

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
	"strconv"
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
	s := html.NewComponents(
		newKeyedTableRow("Total", strconv.Itoa(len(getPlaces(c.document))), true),
	)

	return newCard("Places", noBadgeCount, html.NewTable("", s)).String()
}
