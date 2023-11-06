package html

import (
	"fmt"
	"io"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/gedcom/v39/html/core"
)

// IndividualLink is a hyperlink to an individuals page. The link contains a
// coloured dot to represent their sex and their full name.
type IndividualLink struct {
	individual *gedcom.IndividualNode
	document   *gedcom.Document
	visibility LivingVisibility
	placesMap  map[string]*place
}

func NewIndividualLink(document *gedcom.Document, individual *gedcom.IndividualNode, visibility LivingVisibility, placesMap map[string]*place) *IndividualLink {
	return &IndividualLink{
		individual: individual,
		document:   document,
		visibility: visibility,
		placesMap:  placesMap,
	}
}

func (c *IndividualLink) WriteHTMLTo(w io.Writer) (int64, error) {
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

	dot := core.NewOcticon("primitive-dot", dotStyle)
	individualName := NewIndividualName(c.individual, c.visibility,
		UnknownEmphasis)
	text := core.NewComponents(dot, individualName)

	link := PageIndividual(c.document, c.individual, c.visibility, c.placesMap)

	return core.NewLink(text, link).WriteHTMLTo(w)
}
