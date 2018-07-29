package main

import (
	"github.com/elliotchance/gedcom"
	"strconv"
)

type placeStatistics struct {
	document   *gedcom.Document
}

func newPlaceStatistics(document *gedcom.Document) *placeStatistics {
	return &placeStatistics{
		document:   document,
	}
}

func (c *placeStatistics) String() string {
	s := newComponents(
		newKeyedTableRow("Total", strconv.Itoa(len(getPlaces(c.document))), true),
	)

	return newCard("Places", noBadgeCount, newTable("", s)).String()
}
