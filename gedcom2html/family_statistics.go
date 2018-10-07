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
		n := gedcom.First(gedcom.NodesWithTagPath(family, gedcom.TagMarriage))
		if n != nil {
			marriageEvents += 1
		}

		n = gedcom.First(gedcom.NodesWithTagPath(family, gedcom.TagDivorce))
		if n != nil {
			divorceEvents += 1
		}
	}

	totalFamilies := html.NewNumber(total).String()
	marriageCount := html.NewNumber(marriageEvents).String()
	divorceCount := html.NewNumber(divorceEvents).String()
	totalFamiliesRow := newKeyedTableRow("Total Families", totalFamilies, true)
	marriageCountRow := newKeyedTableRow("Marriage Events", marriageCount, true)
	divorceCountRow := newKeyedTableRow("Divorce Events", divorceCount, true)

	s := html.NewComponents(totalFamiliesRow, marriageCountRow, divorceCountRow)

	return newCard("Families", noBadgeCount, html.NewTable("", s)).String()
}
