package html

import (
	"github.com/elliotchance/gedcom"
	"io"
)

// IndividualEvent is a row in the "Events" section of the individuals page.
type IndividualEvent struct {
	date        string
	place       string
	description Component
	event       gedcom.Node
	individual  *gedcom.IndividualNode
}

func NewIndividualEvent(date, place string, description Component, individual *gedcom.IndividualNode, event gedcom.Node) *IndividualEvent {
	return &IndividualEvent{
		date:        date,
		place:       place,
		description: description,
		individual:  individual,
		event:       event,
	}
}

func (c *IndividualEvent) WriteTo(w io.Writer) (int64, error) {
	kind := c.event.Tag().String()
	placeName := prettyPlaceName(c.place)
	placeLink := NewPlaceLink(c.individual.Document(), placeName)
	age := NewAge(c.individual.AgeAt(c.event))

	return NewTableRow(
		NewTableCell(age).NoWrap(),
		NewTableCell(NewText(kind)).Header(),
		NewTableCell(NewText(c.date)),
		NewTableCell(placeLink),
		NewTableCell(c.description),
	).WriteTo(w)
}
