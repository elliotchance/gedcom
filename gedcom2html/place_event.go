package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
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

	if d := c.node.FirstNodeWithTag(gedcom.TagDate); d != nil {
		date = d.Value()
	}

	for _, individual := range c.document.Individuals() {
		if individual.IsLiving() {
			return ""
		}

		if individual.HasNestedChild(c.node) {
			person = newIndividualLink(c.document, individual).String()
		}
	}

	return fmt.Sprintf(fmt.Sprintf(`
		<tr>
			<td nowrap="nowrap">%s</td>
			<td nowrap="nowrap">%s</td>
			<td nowrap="nowrap">%s</td>
		</tr>`, date, description, person))
}
