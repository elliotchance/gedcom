package html

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
	"io"
)

// IndividualEvent is a row in the "Events" section of the individuals page.
type IndividualEvent struct {
	date        string
	place       string
	description core.Component
	event       gedcom.Node
	individual  *gedcom.IndividualNode
}

func NewIndividualEvent(date, place string, description core.Component, individual *gedcom.IndividualNode, event gedcom.Node) *IndividualEvent {
	return &IndividualEvent{
		date:        date,
		place:       place,
		description: description,
		individual:  individual,
		event:       event,
	}
}

func (c *IndividualEvent) WriteHTMLTo(w io.Writer) (int64, error) {
	kind := c.event.Tag().String()
	placeName := prettyPlaceName(c.place)
	placeLink := NewPlaceLink(c.individual.Document(), placeName)
	age := NewAge(c.individual.AgeAt(c.event))

	return core.NewTableRow(
		core.NewTableCell(age).NoWrap(),
		core.NewTableCell(core.NewText(kind)).Header(),
		core.NewTableCell(core.NewText(c.date)),
		core.NewTableCell(placeLink),
		core.NewTableCell(c.description),
	).WriteHTMLTo(w)
}
