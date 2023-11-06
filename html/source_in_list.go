package html

import (
	"io"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/gedcom/v39/html/core"
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
