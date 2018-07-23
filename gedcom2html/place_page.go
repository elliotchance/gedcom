package main

import (
	"github.com/elliotchance/gedcom"
	"fmt"
)

type placePage struct {
	document *gedcom.Document
	placeKey string
}

func newPlacePage(document *gedcom.Document, placeKey string) *placePage {
	return &placePage{
		document:   document,
		placeKey: placeKey,
	}
}

func (c *placePage) String() string {
	place := getPlaces(c.document)[c.placeKey]

	table := []fmt.Stringer{
		newTableHead("Date", "Event", "Individual"),
	}

	for _, node := range place.nodes {
		table = append(table, newPlaceEvent(c.document, node))
	}

	return newPage(
		place.prettyName,
		newComponents(
			newHeader(c.document, place.prettyName, selectedExtraTab),
			newBigName(place.prettyName),
			newSpace(),
			newRow(
				newColumn(entireRow, newTable("", table...)),
			),
		),
	).String()
}
