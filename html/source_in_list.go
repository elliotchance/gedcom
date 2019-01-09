package html

import (
	"github.com/elliotchance/gedcom"
	"io"
)

type SourceInList struct {
	document *gedcom.Document
	source   *gedcom.SourceNode
}

func NewSourceInList(document *gedcom.Document, source *gedcom.SourceNode) *SourceInList {
	return &SourceInList{
		document: document,
		source:   source,
	}
}

func (c *SourceInList) WriteTo(w io.Writer) (int64, error) {
	return NewTableRow(
		NewTableCell(NewSourceLink(c.source)),
	).WriteTo(w)
}
