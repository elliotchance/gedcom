package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
)

type sourcePage struct {
	document          *gedcom.Document
	source            *gedcom.SourceNode
	googleAnalyticsID string
}

func newSourcePage(document *gedcom.Document, source *gedcom.SourceNode, googleAnalyticsID string) *sourcePage {
	return &sourcePage{
		document:          document,
		source:            source,
		googleAnalyticsID: googleAnalyticsID,
	}
}

func (c *sourcePage) String() string {
	table := []fmt.Stringer{
		html.NewTableHead("Key", "Value"),
	}

	for _, node := range c.source.Nodes() {
		table = append(table, newSourceProperty(c.document, node))
	}

	return html.NewPage(
		c.source.Title(),
		html.NewComponents(
			newHeader(c.document, "Source", selectedExtraTab),
			html.NewBigTitle(c.source.Title()),
			html.NewSpace(),
			html.NewRow(
				html.NewColumn(html.EntireRow, html.NewTable("", table...)),
			),
		),
		c.googleAnalyticsID,
	).String()
}
