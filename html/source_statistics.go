package html

import (
	"io"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/gedcom/v39/html/core"
)

type SourceStatistics struct {
	document *gedcom.Document
}

func NewSourceStatistics(document *gedcom.Document) *SourceStatistics {
	return &SourceStatistics{
		document: document,
	}
}

func (c *SourceStatistics) WriteHTMLTo(w io.Writer) (int64, error) {
	sources := c.document.Sources()
	total := core.NewNumber(len(sources))
	s := core.NewComponents(
		core.NewKeyedTableRow("Total", total, true),
	)

	return core.NewCard(core.NewText("Sources"), core.CardNoBadgeCount,
		core.NewTable("", s)).WriteHTMLTo(w)
}
