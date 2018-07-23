package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"sort"
)

// placeListPage lists all places.
type placeListPage struct {
	document *gedcom.Document
}

func newPlaceListPage(document *gedcom.Document) *placeListPage {
	return &placeListPage{
		document: document,
	}
}

func (c *placeListPage) String() string {
	table := []fmt.Stringer{
		newTableHead("Name"),
	}

	places := getPlaces(c.document)

	sortedPlaces := []*place{}
	for _, placeName := range places {
		sortedPlaces = append(sortedPlaces, placeName)
	}

	sort.SliceStable(sortedPlaces, func(i, j int) bool {
		return sortedPlaces[i].prettyName < sortedPlaces[j].prettyName
	})

	for _, place := range sortedPlaces {
		table = append(table, newPlaceInList(c.document, place))
	}

	return newPage("Places", newComponents(
		newHeader(c.document, "", selectedPlacesTab),
		newRow(
			newColumn(entireRow, newTable("", table...)),
		),
	)).String()
}
