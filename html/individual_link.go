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
	visibility LivingVisibility
}

func NewIndividualLink(document *gedcom.Document, individual *gedcom.IndividualNode, visibility LivingVisibility) *IndividualLink {
	return &IndividualLink{
		individual: individual,
		document:   document,
		visibility: visibility,
	}
}

func (c *IndividualLink) WriteTo(w io.Writer) (int64, error) {
	if c.individual.IsLiving() {
		switch c.visibility {
		case LivingVisibilityHide:
			return writeNothing()

		case LivingVisibilityShow, LivingVisibilityPlaceholder:
			// Proceed.
		}
	}

	dotColor := colorForIndividual(c.individual)
	dotStyle := fmt.Sprintf("color: %s; font-size: 18px", dotColor)

	dot := NewOcticon("primitive-dot", dotStyle)
	individualName := NewIndividualName(c.individual, c.visibility,
		UnknownEmphasis)
	text := NewComponents(dot, individualName)

	link := PageIndividual(c.document, c.individual, c.visibility)

	return NewLink(text, link).WriteTo(w)
}
