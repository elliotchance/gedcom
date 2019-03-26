package html

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
	"io"
)

type SourcePage struct {
	document          *gedcom.Document
	source            *gedcom.SourceNode
	googleAnalyticsID string
	options           *PublishShowOptions
	indexLetters      []rune
}

func NewSourcePage(document *gedcom.Document, source *gedcom.SourceNode, googleAnalyticsID string, options *PublishShowOptions, indexLetters []rune) *SourcePage {
	return &SourcePage{
		document:          document,
		source:            source,
		googleAnalyticsID: googleAnalyticsID,
		options:           options,
		indexLetters:      indexLetters,
	}
}

func (c *SourcePage) WriteHTMLTo(w io.Writer) (int64, error) {
	table := []core.Component{
		core.NewTableHead("Key", "Value"),
	}

	for _, node := range c.source.Nodes() {
		table = append(table, NewSourceProperty(c.document, node))
	}

	return core.NewPage(
		c.source.Title(),
		core.NewComponents(
			NewPublishHeader(c.document, "Source", selectedExtraTab, c.options,
				c.indexLetters),
			core.NewBigTitle(1, core.NewText(c.source.Title())),
			core.NewSpace(),
			core.NewRow(
				core.NewColumn(core.EntireRow, core.NewTable("", table...)),
			),
		),
		c.googleAnalyticsID,
	).WriteHTMLTo(w)
}
