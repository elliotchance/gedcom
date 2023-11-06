package html

import (
	"io"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/gedcom/v39/html/core"
)

type SourcePage struct {
	document          *gedcom.Document
	source            *gedcom.SourceNode
	googleAnalyticsID string
	options           *PublishShowOptions
	indexLetters      []rune
	placesMap         map[string]*place
}

func NewSourcePage(document *gedcom.Document, source *gedcom.SourceNode, googleAnalyticsID string, options *PublishShowOptions, indexLetters []rune, placesMap map[string]*place) *SourcePage {
	return &SourcePage{
		document:          document,
		source:            source,
		googleAnalyticsID: googleAnalyticsID,
		options:           options,
		indexLetters:      indexLetters,
		placesMap:         placesMap,
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
			NewPublishHeader(c.document, "Source", selectedExtraTab,
				c.options, c.indexLetters, c.placesMap),
			core.NewBigTitle(1, core.NewText(c.source.Title())),
			core.NewSpace(),
			core.NewRow(
				core.NewColumn(core.EntireRow, core.NewTable("", table...)),
			),
		),
		c.googleAnalyticsID,
	).WriteHTMLTo(w)
}
