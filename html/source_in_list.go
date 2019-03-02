package html

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
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

func (c *SourceInList) WriteHTMLTo(w io.Writer) (int64, error) {
	return core.NewTableRow(
		core.NewTableCell(NewSourceLink(c.source)),
	).WriteHTMLTo(w)
}
