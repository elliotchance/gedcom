package main

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
)

type familyStatistics struct {
	document *gedcom.Document
}

func newFamilyStatistics(document *gedcom.Document) *familyStatistics {
	return &familyStatistics{
		document: document,
	}
}

func (c *familyStatistics) String() string {
	total := len(c.document.Families())
	marriageEvents := 0
	divorceEvents := 0

	for _, family := range c.document.Families() {
		if n := gedcom.First(gedcom.NodesWithTagPath(family, gedcom.TagMarriage)); n != nil {
			marriageEvents += 1
		}

		if n := gedcom.First(gedcom.NodesWithTagPath(family, gedcom.TagDivorce)); n != nil {
			divorceEvents += 1
		}
	}

	s := html.NewComponents(
		newKeyedTableRow("Total Families", html.NewNumber(total).String(), true),
		newKeyedTableRow("Marriage Events", html.NewNumber(marriageEvents).String(), true),
		newKeyedTableRow("Divorce Events", html.NewNumber(divorceEvents).String(), true),
	)

	return newCard("Families", noBadgeCount, html.NewTable("", s)).String()
}
