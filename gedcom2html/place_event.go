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

	if d := gedcom.First(gedcom.NodesWithTag(c.node, gedcom.TagDate)); d != nil {
		date = d.Value()
	}

	for _, individual := range c.document.Individuals() {
		if gedcom.HasNestedNode(individual, c.node) {
			if individual.IsLiving() {
				return ""
			}

			person = newIndividualLink(c.document, individual).String()
			break
		}
	}

	return html.Sprintf(`
		<tr>
			<td nowrap="nowrap">%s</td>
			<td nowrap="nowrap">%s</td>
			<td nowrap="nowrap">%s</td>
		</tr>`, date, description, person)
}
