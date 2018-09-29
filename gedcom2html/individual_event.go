package main

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
)

// individualEvent is a row in the "Events" section of the individuals page.
type individualEvent struct {
	kind        string
	date        string
	place       string
	description string
	document    *gedcom.Document
}

func newIndividualEvent(kind, date, place, description string, document *gedcom.Document) *individualEvent {
	return &individualEvent{
		kind:        kind,
		date:        date,
		place:       place,
		description: description,
		document:    document,
	}
}

func (c *individualEvent) String() string {
	placeLink := newPlaceLink(c.document, c.place)

	return html.Sprintf(`
		<tr>
			<th>%s</th>
			<td>%s</td>
			<td>%s</td>
			<td>%s</td>
		</tr>`, c.kind, c.date, placeLink.String(), c.description)
}
