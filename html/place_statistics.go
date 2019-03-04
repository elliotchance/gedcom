package html

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
	"io"
)

type PlaceStatistics struct {
	document *gedcom.Document
}

func newPlaceStatistics(document *gedcom.Document) *PlaceStatistics {
	return &PlaceStatistics{
		document: document,
	}
}

func (c *PlaceStatistics) WriteHTMLTo(w io.Writer) (int64, error) {
	places := GetPlaces(c.document)
	total := core.NewNumber(len(places))
	s := core.NewComponents(
		core.NewKeyedTableRow("Total", total, true),
	)

	return core.NewCard(core.NewText("Places"), core.CardNoBadgeCount,
		core.NewTable("", s)).WriteHTMLTo(w)
}
