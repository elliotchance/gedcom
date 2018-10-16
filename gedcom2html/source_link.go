package main

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
)

type sourceLink struct {
	source *gedcom.SourceNode
}

func newSourceLink(source *gedcom.SourceNode) *sourceLink {
	return &sourceLink{
		source: source,
	}
}

func (c *sourceLink) String() string {
	text := c.source.Title()
	destination := pageSource(c.source)

	return html.NewLink(text, destination).String()
}
