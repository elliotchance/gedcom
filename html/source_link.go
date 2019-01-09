package html

import (
	"github.com/elliotchance/gedcom"
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

func (c *SourceLink) WriteTo(w io.Writer) (int64, error) {
	text := c.source.Title()
	destination := PageSource(c.source)

	return NewLink(NewText(text), destination).WriteTo(w)
}
