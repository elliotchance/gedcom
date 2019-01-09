package html

import (
	"github.com/elliotchance/gedcom"
	"io"
)

type SourceListPage struct {
	document          *gedcom.Document
	googleAnalyticsID string
	options           PublishShowOptions
}

func NewSourceListPage(document *gedcom.Document, googleAnalyticsID string, options PublishShowOptions) *SourceListPage {
	return &SourceListPage{
		document:          document,
		googleAnalyticsID: googleAnalyticsID,
		options:           options,
	}
}

func (c *SourceListPage) WriteTo(w io.Writer) (int64, error) {
	table := []Component{
		NewTableHead("Name"),
	}

	for _, source := range c.document.Sources() {
		table = append(table, NewSourceInList(c.document, source))
	}

	return NewPage("Sources", NewComponents(
		NewPublishHeader(c.document, "", selectedSourcesTab, c.options),
		NewRow(
			NewColumn(EntireRow, NewTable("", table...)),
		),
	), c.googleAnalyticsID).WriteTo(w)
}
