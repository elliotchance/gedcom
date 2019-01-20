package html

import (
	"github.com/elliotchance/gedcom"
	"io"
)

type FamilyListPage struct {
	document          *gedcom.Document
	googleAnalyticsID string
	options           PublishShowOptions
	visibility        LivingVisibility
}

func NewFamilyListPage(document *gedcom.Document, googleAnalyticsID string, options PublishShowOptions, visibility LivingVisibility) *FamilyListPage {
	return &FamilyListPage{
		document:          document,
		googleAnalyticsID: googleAnalyticsID,
		options:           options,
		visibility:        visibility,
	}
}

func (c *FamilyListPage) WriteTo(w io.Writer) (int64, error) {
	table := []Component{
		NewTableHead("Husband", "Date", "Wife"),
	}

	for _, family := range c.document.Families() {
		familyInList := NewFamilyInList(c.document, family, c.visibility)
		table = append(table, familyInList)
	}

	column := NewColumn(EntireRow, NewTable("", table...))
	header := NewPublishHeader(c.document, "", selectedFamiliesTab, c.options)
	components := NewComponents(header, NewRow(column))

	return NewPage("Families", components, c.googleAnalyticsID).WriteTo(w)
}
