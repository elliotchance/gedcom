package html

import (
	"io"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/gedcom/v39/html/core"
)

type PlaceInList struct {
	document  *gedcom.Document
	place     *place
	placesMap map[string]*place
}

func NewPlaceInList(document *gedcom.Document, place *place, placesMap map[string]*place) *PlaceInList {
	return &PlaceInList{
		document:  document,
		place:     place,
		placesMap: placesMap,
	}
}

func (c *PlaceInList) WriteHTMLTo(w io.Writer) (int64, error) {
	placeLink := NewPlaceLink(c.document, c.place.PrettyName, c.placesMap)
	countBadge := core.NewCountBadge(len(c.place.nodes))
	content := core.NewComponents(placeLink, countBadge)

	return core.NewTableRow(
		core.NewTableCell(content).NoWrap(),
	).WriteHTMLTo(w)
}
