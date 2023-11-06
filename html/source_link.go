package html

import (
	"io"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/gedcom/v39/html/core"
)

type SourceLink struct {
	source *gedcom.SourceNode
}

func NewSourceLink(source *gedcom.SourceNode) *SourceLink {
	return &SourceLink{
		source: source,
	}
}

func (c *SourceLink) WriteHTMLTo(w io.Writer) (int64, error) {
	text := c.source.Title()
	destination := PageSource(c.source)

	return core.NewLink(core.NewText(text), destination).WriteHTMLTo(w)
}
