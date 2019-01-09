package html

import (
	"github.com/elliotchance/gedcom"
	"io"
)

type PlaceEvent struct {
	node     gedcom.Node
	document *gedcom.Document
}

func NewPlaceEvent(document *gedcom.Document, node gedcom.Node) *PlaceEvent {
	return &PlaceEvent{
		document: document,
		node:     node,
	}
}

func (c *PlaceEvent) WriteTo(w io.Writer) (int64, error) {
	date := ""
	description := c.node.Tag().String()
	var person Component = NewEmpty()

	d := gedcom.Dates(c.node).Minimum()

	if d != nil {
		date = d.Value()
	}

	individual := individualForNode(c.node)
	if individual != nil && !individual.IsLiving() {
		person = NewIndividualLink(c.document, individual)
	}

	return NewTableRow(
		NewTableCell(NewText(date)).NoWrap(),
		NewTableCell(NewText(description)).NoWrap(),
		NewTableCell(person).NoWrap(),
	).WriteTo(w)
}
