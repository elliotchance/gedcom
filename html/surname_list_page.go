package html

import (
	"github.com/elliotchance/gedcom"
	"io"
)

type SurnameListPage struct {
	document          *gedcom.Document
	googleAnalyticsID string
	options           PublishShowOptions
}

func NewSurnameListPage(document *gedcom.Document, googleAnalyticsID string, options PublishShowOptions) *SurnameListPage {
	return &SurnameListPage{
		document:          document,
		googleAnalyticsID: googleAnalyticsID,
		options:           options,
	}
}

func (c *SurnameListPage) WriteTo(w io.Writer) (int64, error) {
	table := []Component{
		NewTableHead("Surname", "Number of Individuals"),
	}

	for _, surname := range getSurnames(c.document) {
		table = append(table, NewSurnameInList(c.document, surname))
	}

	return NewPage("Surnames", NewComponents(
		NewPublishHeader(c.document, "", selectedSurnamesTab, c.options),
		NewRow(
			NewColumn(EntireRow, NewTable("", table...)),
		),
	), c.googleAnalyticsID).WriteTo(w)
}
