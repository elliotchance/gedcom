package main

import (
	"github.com/elliotchance/gedcom"
	"fmt"
)

type sourcePage struct {
	document *gedcom.Document
	source   *gedcom.SourceNode
}

func newSourcePage(document *gedcom.Document, source *gedcom.SourceNode) *sourcePage {
	return &sourcePage{
		document: document,
		source:   source,
	}
}

func (c *sourcePage) String() string {
	table := []fmt.Stringer{
		newTableHead("Key", "Value"),
	}

	for _, node := range c.source.Nodes() {
		table = append(table, newSourceProperty(c.document, node))
	}

	return newPage(
		c.source.Title(),
		newComponents(
			newHeader(c.document, "Source", selectedExtraTab),
			newBigName(c.source.Title()),
			newSpace(),
			newRow(
				newColumn(entireRow, newTable("", table...)),
			),
		),
	).String()
}
