package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
)

type sourceLink struct {
	source   *gedcom.SourceNode
}

func newSourceLink(source *gedcom.SourceNode) *sourceLink {
	return &sourceLink{
		source:   source,
	}
}

func (c *sourceLink) String() string {
	return fmt.Sprintf(`
		<a href="%s">%s</a>`,
		pageSource(c.source), c.source.Title())
}
