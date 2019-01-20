package html

import (
	"github.com/elliotchance/gedcom"
	"io"
)

type PlaceStatistics struct {
	document *gedcom.Document
}

func newPlaceStatistics(document *gedcom.Document) *PlaceStatistics {
	return &PlaceStatistics{
		document: document,
	}
}

func (c *PlaceStatistics) WriteTo(w io.Writer) (int64, error) {
	places := GetPlaces(c.document)
	total := NewNumber(len(places))
	s := NewComponents(
		NewKeyedTableRow("Total", total, true),
	)

	return NewCard("Places", noBadgeCount, NewTable("", s)).WriteTo(w)
}
