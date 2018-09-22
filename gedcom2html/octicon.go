package main

import (
	"github.com/elliotchance/gedcom/html"
)

type octicon struct {
	name  string
	style string
}

func newOcticon(name, style string) *octicon {
	return &octicon{
		name:  name,
		style: style,
	}
}

func (c *octicon) String() string {
	return html.Sprintf(`<span class="octicon octicon-%s" style="%s"></span>`,
		c.name, c.style)
}
