package main

import (
	"github.com/elliotchance/gedcom"
	"strconv"
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

	s := newComponents(
		newKeyedTableRow("Total Families", strconv.Itoa(total), true),
		newKeyedTableRow("Marriage Events", strconv.Itoa(marriageEvents), true),
		newKeyedTableRow("Divorce Events", strconv.Itoa(divorceEvents), true),
	)

	return newCard("Families", noBadgeCount, newTable("", s)).String()
}
