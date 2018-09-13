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
	return html.Sprintf(`
		<tr>
			<td nowrap="nowrap">%s %s</td>
		</tr>`,
		newPlaceLink(c.document, c.place.prettyName),
		newCountBadge(len(c.place.nodes)))
}
