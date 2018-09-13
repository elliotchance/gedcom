package main

import (
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
	return html.Sprintf(`
		<span class="octicon octicon-primitive-dot" style="color: %s; font-size: 18px"></span>
		<a href="%s">%s</a>`,
		colorForIndividual(c.individual),
		pageIndividual(c.document, c.individual),
		html.NewIndividualName(c.individual, false, html.UnknownEmphasis))
}
