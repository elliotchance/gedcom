package html

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
	"io"
)

type PlaceEvent struct {
	node       gedcom.Node
	document   *gedcom.Document
	visibility LivingVisibility
	placesMap  map[string]*place
}

func NewPlaceEvent(document *gedcom.Document, node gedcom.Node, visibility LivingVisibility, placesMap map[string]*place) *PlaceEvent {
	return &PlaceEvent{
		document:   document,
		node:       node,
		visibility: visibility,
		placesMap:  placesMap,
	}
}

func (c *PlaceEvent) WriteHTMLTo(w io.Writer) (int64, error) {
	date := ""
	description := c.node.Tag().String()

	d := gedcom.Dates(c.node).Minimum()

	if d != nil {
		date = d.Value()
	}

	individual := individualForNode(c.document, c.node)
	var person core.Component = NewIndividualLink(c.document, individual,
		c.visibility, c.placesMap)
	isLiving := individual != nil && individual.IsLiving()

	if isLiving {
		switch c.visibility {
		case LivingVisibilityHide:
			return writeNothing()

		case LivingVisibilityShow:
			// Proceed.

		case LivingVisibilityPlaceholder:
			person = core.NewEmpty()
		}
	}

	return core.NewTableRow(
		core.NewTableCell(core.NewText(date)).NoWrap(),
		core.NewTableCell(core.NewText(description)).NoWrap(),
		core.NewTableCell(person).NoWrap(),
	).WriteHTMLTo(w)
}
