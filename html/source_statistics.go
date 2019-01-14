package html

import (
	"github.com/elliotchance/gedcom"
	"io"
)

type SourceStatistics struct {
	document *gedcom.Document
}

func NewSourceStatistics(document *gedcom.Document) *SourceStatistics {
	return &SourceStatistics{
		document: document,
	}
}

func (c *SourceStatistics) WriteTo(w io.Writer) (int64, error) {
	sources := c.document.Sources()
	total := NewNumber(len(sources))
	s := NewComponents(
		NewKeyedTableRow("Total", total, true),
	)

	return NewCard(NewText("Sources"), noBadgeCount, NewTable("", s)).WriteTo(w)
}
