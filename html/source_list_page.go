package html

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
	"io"
)

type SourceListPage struct {
	document          *gedcom.Document
	googleAnalyticsID string
	options           *PublishShowOptions
	indexLetters      []rune
}

func NewSourceListPage(document *gedcom.Document, googleAnalyticsID string, options *PublishShowOptions, indexLetters []rune) *SourceListPage {
	return &SourceListPage{
		document:          document,
		googleAnalyticsID: googleAnalyticsID,
		options:           options,
		indexLetters:      indexLetters,
	}
}

func (c *SourceListPage) WriteHTMLTo(w io.Writer) (int64, error) {
	table := []core.Component{
		core.NewTableHead("Name"),
	}

	for _, source := range c.document.Sources() {
		table = append(table, NewSourceInList(c.document, source))
	}

	return core.NewPage("Sources", core.NewComponents(
		NewPublishHeader(c.document, "", selectedSourcesTab, c.options, c.indexLetters),
		core.NewRow(
			core.NewColumn(core.EntireRow, core.NewTable("", table...)),
		),
	), c.googleAnalyticsID).WriteHTMLTo(w)
}
