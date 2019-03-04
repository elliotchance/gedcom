package html

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
	"io"
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
