package html

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
	"io"
)

// AllParentButtons represent one or more families that an individual belongs
// to. These are show as large buttons above the large name of the person in on
// their individual page.
type AllParentButtons struct {
	individual *gedcom.IndividualNode
	document   *gedcom.Document
	visibility LivingVisibility
	placesMap  map[string]*place
}

func NewAllParentButtons(document *gedcom.Document, individual *gedcom.IndividualNode, visibility LivingVisibility, placesMap map[string]*place) *AllParentButtons {
	return &AllParentButtons{
		individual: individual,
		document:   document,
		visibility: visibility,
		placesMap:  placesMap,
	}
}

func (c *AllParentButtons) WriteHTMLTo(w io.Writer) (int64, error) {
	families := c.individual.Families()
	components := []core.Component{}

	for _, family := range families {
		husbandMatches := family.Husband().IsIndividual(c.individual)
		wifeMatches := family.Wife().IsIndividual(c.individual)

		if husbandMatches || wifeMatches {
			continue
		}

		components = append(components,
			NewParentButtons(c.document, family, c.visibility, c.placesMap))
	}

	// If there are no families we still want to show an empty family. We just
	// create a dummy family that has no child nodes.
	if len(components) == 0 {
		familyNode := gedcom.NewDocument().AddFamily("")
		components = []core.Component{
			NewParentButtons(c.document, familyNode, c.visibility, c.placesMap),
		}
	}

	return core.NewComponents(components...).WriteHTMLTo(w)
}
