package html

import (
	"github.com/elliotchance/gedcom"
	"io"
)

type PlaceEvent struct {
	node       gedcom.Node
	document   *gedcom.Document
	visibility LivingVisibility
}

func NewPlaceEvent(document *gedcom.Document, node gedcom.Node, visibility LivingVisibility) *PlaceEvent {
	return &PlaceEvent{
		document:   document,
		node:       node,
		visibility: visibility,
	}
}

func (c *PlaceEvent) WriteTo(w io.Writer) (int64, error) {
	date := ""
	description := c.node.Tag().String()

	d := gedcom.Dates(c.node).Minimum()

	if d != nil {
		date = d.Value()
	}

	individual := individualForNode(c.node)
	var person Component = NewIndividualLink(c.document, individual, c.visibility)
	isLiving := individual != nil && individual.IsLiving()

	if isLiving {
		switch c.visibility {
		case LivingVisibilityHide:
			return writeNothing()

		case LivingVisibilityShow:
			// Proceed.

		case LivingVisibilityPlaceholder:
			person = NewEmpty()
		}
	}

	return NewTableRow(
		NewTableCell(NewText(date)).NoWrap(),
		NewTableCell(NewText(description)).NoWrap(),
		NewTableCell(person).NoWrap(),
	).WriteTo(w)
}
