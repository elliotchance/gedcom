package html

import (
	"github.com/elliotchance/gedcom"
	"io"
)

type PlaceInList struct {
	document *gedcom.Document
	place    *place
}

func NewPlaceInList(document *gedcom.Document, place *place) *PlaceInList {
	return &PlaceInList{
		document: document,
		place:    place,
	}
}

func (c *PlaceInList) WriteTo(w io.Writer) (int64, error) {
	placeLink := NewPlaceLink(c.document, c.place.PrettyName)
	countBadge := NewCountBadge(len(c.place.nodes))
	content := NewComponents(placeLink, countBadge)

	return NewTableRow(
		NewTableCell(content).NoWrap(),
	).WriteTo(w)
}
