package html

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
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

func (c *PlaceInList) WriteHTMLTo(w io.Writer) (int64, error) {
	placeLink := NewPlaceLink(c.document, c.place.PrettyName)
	countBadge := core.NewCountBadge(len(c.place.nodes))
	content := core.NewComponents(placeLink, countBadge)

	return core.NewTableRow(
		core.NewTableCell(content).NoWrap(),
	).WriteHTMLTo(w)
}
