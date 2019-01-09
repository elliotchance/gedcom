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
	total := NewNumber(len(c.document.Sources()))
	s := NewComponents(
		NewKeyedTableRow("Total", total, true),
	)

	return NewCard("Sources", noBadgeCount, NewTable("", s)).WriteTo(w)
}
