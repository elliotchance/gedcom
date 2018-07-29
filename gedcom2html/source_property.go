package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
)

type sourceProperty struct {
	document *gedcom.Document
	node     gedcom.Node
}

func newSourceProperty(document *gedcom.Document, node gedcom.Node) *sourceProperty {
	return &sourceProperty{
		document: document,
		node:     node,
	}
}

func (c *sourceProperty) String() string {
	s := fmt.Sprintf(`
		<tr>
			<th nowrap="nowrap">%s</th>
			<td>%s</td>
		</tr>`, c.node.Tag().String(), c.node.Value())

	for _, node := range c.node.Nodes() {
		s += newSourceProperty(c.document, node).String()
	}

	return s
}
