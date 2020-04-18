package html

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
	"io"
)

type PlaceStatistics struct {
	document  *gedcom.Document
	placesMap map[string]*place
}

func newPlaceStatistics(document *gedcom.Document, placesMap map[string]*place) *PlaceStatistics {
	return &PlaceStatistics{
		document:  document,
		placesMap: placesMap,
	}
}

func (c *PlaceStatistics) WriteHTMLTo(w io.Writer) (int64, error) {
	total := core.NewNumber(len(c.placesMap))
	s := core.NewComponents(
		core.NewKeyedTableRow("Total", total, true),
	)

	return core.NewCard(core.NewText("Places"), core.CardNoBadgeCount,
		core.NewTable("", s)).WriteHTMLTo(w)
}
