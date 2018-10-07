package main

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
)

type placeInList struct {
	document *gedcom.Document
	place    *place
}

func newPlaceInList(document *gedcom.Document, place *place) *placeInList {
	return &placeInList{
		document: document,
		place:    place,
	}
}

func (c *placeInList) String() string {
	placeLink := newPlaceLink(c.document, c.place.prettyName)
	countBadge := newCountBadge(len(c.place.nodes))
	content := html.NewComponents(placeLink, countBadge)

	return html.NewTableRow(
		html.NewTableCell(content).NoWrap(),
	).String()
}
