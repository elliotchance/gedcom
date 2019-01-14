package html

import (
	"github.com/elliotchance/gedcom"
	"io"
)

type FamilyStatistics struct {
	document *gedcom.Document
}

func NewFamilyStatistics(document *gedcom.Document) *FamilyStatistics {
	return &FamilyStatistics{
		document: document,
	}
}

func (c *FamilyStatistics) WriteTo(w io.Writer) (int64, error) {
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

	totalFamilies := NewNumber(total)
	marriageCount := NewNumber(marriageEvents)
	divorceCount := NewNumber(divorceEvents)
	totalFamiliesRow := NewKeyedTableRow("Total Families", totalFamilies, true)
	marriageCountRow := NewKeyedTableRow("Marriage Events", marriageCount, true)
	divorceCountRow := NewKeyedTableRow("Divorce Events", divorceCount, true)

	s := NewComponents(totalFamiliesRow, marriageCountRow, divorceCountRow)

	return NewCard(NewText("Families"), noBadgeCount, NewTable("", s)).
		WriteTo(w)
}
