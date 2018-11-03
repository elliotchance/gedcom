package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
)

// individualEvent is a row in the "Events" section of the individuals page.
type individualEvent struct {
	date        string
	place       string
	description fmt.Stringer
	event       gedcom.Node
	individual  *gedcom.IndividualNode
}

func newIndividualEvent(date, place string, description fmt.Stringer, individual *gedcom.IndividualNode, event gedcom.Node) *individualEvent {
	return &individualEvent{
		date:        date,
		place:       place,
		description: description,
		individual:  individual,
		event:       event,
	}
}

func (c *individualEvent) String() string {
	kind := c.event.Tag().String()
	placeLink := newPlaceLink(c.individual.Document(), c.place)
	age := html.NewAge(c.individual.AgeAt(c.event))

	return html.NewTableRow(
		html.NewTableCell(age).NoWrap(),
		html.NewTableCell(html.NewText(kind)).Header(),
		html.NewTableCell(html.NewText(c.date)),
		html.NewTableCell(placeLink),
		html.NewTableCell(c.description),
	).String()
}
