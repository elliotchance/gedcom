package html

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"io"
)

// IndividualLink is a hyperlink to an individuals page. The link contains a
// coloured dot to represent their sex and their full name.
type IndividualLink struct {
	individual *gedcom.IndividualNode
	document   *gedcom.Document
}

func NewIndividualLink(document *gedcom.Document, individual *gedcom.IndividualNode) *IndividualLink {
	return &IndividualLink{
		individual: individual,
		document:   document,
	}
}

func (c *IndividualLink) WriteTo(w io.Writer) (int64, error) {
	dotColor := colorForIndividual(c.individual)
	dotStyle := fmt.Sprintf("color: %s; font-size: 18px", dotColor)

	dot := NewOcticon("primitive-dot", dotStyle)
	individualName := NewIndividualName(c.individual, false,
		UnknownEmphasis)
	text := NewComponents(dot, individualName)

	link := PageIndividual(c.document, c.individual)

	return NewLink(text, link).WriteTo(w)
}
