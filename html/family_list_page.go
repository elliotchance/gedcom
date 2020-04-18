package html

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
	"io"
)

type FamilyListPage struct {
	document          *gedcom.Document
	googleAnalyticsID string
	options           *PublishShowOptions
	indexLetters      []rune
	placesMap         map[string]*place
}

func NewFamilyListPage(document *gedcom.Document, googleAnalyticsID string, options *PublishShowOptions, indexLetters []rune, placesMap map[string]*place) *FamilyListPage {
	return &FamilyListPage{
		document:          document,
		googleAnalyticsID: googleAnalyticsID,
		options:           options,
		indexLetters:      indexLetters,
		placesMap:         placesMap,
	}
}

func (c *FamilyListPage) WriteHTMLTo(w io.Writer) (int64, error) {
	table := []core.Component{
		core.NewTableHead("Husband", "Date", "Wife"),
	}

	for _, family := range c.document.Families() {
		familyInList := NewFamilyInList(c.document, family,
			c.options.LivingVisibility, c.placesMap)
		table = append(table, familyInList)
	}

	column := core.NewColumn(core.EntireRow, core.NewTable("", table...))
	header := NewPublishHeader(c.document, "", selectedFamiliesTab,
		c.options, c.indexLetters, c.placesMap)
	components := core.NewComponents(header, core.NewRow(column))

	return core.NewPage("Families", components, c.googleAnalyticsID).
		WriteHTMLTo(w)
}
