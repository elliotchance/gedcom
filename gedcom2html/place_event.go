package main

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
)

type placeEvent struct {
	node     gedcom.Node
	document *gedcom.Document
}

func newPlaceEvent(document *gedcom.Document, node gedcom.Node) *placeEvent {
	return &placeEvent{
		document: document,
		node:     node,
	}
}

func (c *placeEvent) String() string {
	date := ""
	description := c.node.Tag().String()
	person := ""

	d := gedcom.MinimumDateNode(gedcom.Dates(c.node))

	if d != nil {
		date = d.Value()
	}

	individual := individualForNode(c.node)
	if individual != nil && !individual.IsLiving() {
		person = newIndividualLink(c.document, individual).String()
	}

	return html.Sprintf(`
		<tr>
			<td nowrap="nowrap">%s</td>
			<td nowrap="nowrap">%s</td>
			<td nowrap="nowrap">%s</td>
		</tr>`, date, description, person)
}
