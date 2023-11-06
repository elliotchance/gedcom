package html

import (
	"io"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/gedcom/v39/html/core"
)

type SurnameListPage struct {
	document          *gedcom.Document
	googleAnalyticsID string
	options           *PublishShowOptions
	indexLetters      []rune
	placesMap         map[string]*place
}

func NewSurnameListPage(document *gedcom.Document, googleAnalyticsID string, options *PublishShowOptions, indexLetters []rune, placesMap map[string]*place) *SurnameListPage {
	return &SurnameListPage{
		document:          document,
		googleAnalyticsID: googleAnalyticsID,
		options:           options,
		indexLetters:      indexLetters,
		placesMap:         placesMap,
	}
}

func (c *SurnameListPage) WriteHTMLTo(w io.Writer) (int64, error) {
	table := []core.Component{
		core.NewTableHead("Surname", "Number of Individuals"),
	}

	for _, surname := range getSurnames(c.document).Strings() {
		table = append(table, NewSurnameInList(c.document, surname))
	}

	return core.NewPage("Surnames", core.NewComponents(
		NewPublishHeader(c.document, "", selectedSurnamesTab, c.options,
			c.indexLetters, c.placesMap),
		core.NewRow(
			core.NewColumn(core.EntireRow, core.NewTable("", table...)),
		),
	), c.googleAnalyticsID).WriteHTMLTo(w)
}
