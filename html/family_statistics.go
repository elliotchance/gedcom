package html

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
	"github.com/elliotchance/gedcom/tag"
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

func (c *FamilyStatistics) WriteHTMLTo(w io.Writer) (int64, error) {
	total := len(c.document.Families())
	marriageEvents := 0
	divorceEvents := 0

	for _, family := range c.document.Families() {
		n := gedcom.First(gedcom.NodesWithTagPath(family, tag.TagMarriage))
		if n != nil {
			marriageEvents += 1
		}

		n = gedcom.First(gedcom.NodesWithTagPath(family, tag.TagDivorce))
		if n != nil {
			divorceEvents += 1
		}
	}

	totalFamilies := core.NewNumber(total)
	marriageCount := core.NewNumber(marriageEvents)
	divorceCount := core.NewNumber(divorceEvents)
	totalFamiliesRow := core.NewKeyedTableRow("Total Families", totalFamilies, true)
	marriageCountRow := core.NewKeyedTableRow("Marriage Events", marriageCount, true)
	divorceCountRow := core.NewKeyedTableRow("Divorce Events", divorceCount, true)

	s := core.NewComponents(totalFamiliesRow, marriageCountRow, divorceCountRow)

	return core.NewCard(core.NewText("Families"), core.CardNoBadgeCount,
		core.NewTable("", s)).WriteHTMLTo(w)
}
