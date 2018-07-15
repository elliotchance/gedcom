package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
)

// individualLink is a hyperlink to an individuals page. The link contains a
// coloured dot to represent their sex and their full name.
type individualLink struct {
	individual *gedcom.IndividualNode
}

func newIndividualLink(individual *gedcom.IndividualNode) *individualLink {
	return &individualLink{
		individual: individual,
	}
}

func (c *individualLink) String() string {
	return fmt.Sprintf(`
		<span class="octicon octicon-primitive-dot" style="color: %s; font-size: 18px"></span>
		<a href="%s">%s</a>`,
		colorForIndividual(c.individual),
		pageIndividual(c.individual), newIndividualName(c.individual))
}
