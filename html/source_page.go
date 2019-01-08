package html

import (
	"github.com/elliotchance/gedcom"
	"io"
)

type SourcePage struct {
	document          *gedcom.Document
	source            *gedcom.SourceNode
	googleAnalyticsID string
	options           PublishShowOptions
}

func NewSourcePage(document *gedcom.Document, source *gedcom.SourceNode, googleAnalyticsID string, options PublishShowOptions) *SourcePage {
	return &SourcePage{
		document:          document,
		source:            source,
		googleAnalyticsID: googleAnalyticsID,
		options:           options,
	}
}

func (c *SourcePage) WriteTo(w io.Writer) (int64, error) {
	table := []Component{
		NewTableHead("Key", "Value"),
	}

	for _, node := range c.source.Nodes() {
		table = append(table, NewSourceProperty(c.document, node))
	}

	return NewPage(
		c.source.Title(),
		NewComponents(
			NewPublishHeader(c.document, "Source", selectedExtraTab, c.options),
			NewBigTitle(1, NewText(c.source.Title())),
			NewSpace(),
			NewRow(
				NewColumn(EntireRow, NewTable("", table...)),
			),
		),
		c.googleAnalyticsID,
	).WriteTo(w)
}
