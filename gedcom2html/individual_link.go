package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
)

// individualLink is a hyperlink to an individuals page. The link contains a
// coloured dot to represent their sex and their full name.
type individualLink struct {
	individual *gedcom.IndividualNode
	document   *gedcom.Document
}

func newIndividualLink(document *gedcom.Document, individual *gedcom.IndividualNode) *individualLink {
	return &individualLink{
		individual: individual,
		document:   document,
	}
}

func (c *individualLink) String() string {
	dotColor := colorForIndividual(c.individual)
	dotStyle := fmt.Sprintf("color: %s; font-size: 18px", dotColor)

	return html.Sprintf(`<a href="%s">%s%s</a>`,
		pageIndividual(c.document, c.individual),
		newOcticon("primitive-dot", dotStyle).String(),
		html.NewIndividualName(c.individual, false, html.UnknownEmphasis))
}
