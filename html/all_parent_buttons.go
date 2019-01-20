package html

import (
	"github.com/elliotchance/gedcom"
	"io"
)

// AllParentButtons represent one or more families that an individual belongs
// to. These are show as large buttons above the large name of the person in on
// their individual page.
type AllParentButtons struct {
	individual *gedcom.IndividualNode
	document   *gedcom.Document
	visibility LivingVisibility
}

func NewAllParentButtons(document *gedcom.Document, individual *gedcom.IndividualNode, visibility LivingVisibility) *AllParentButtons {
	return &AllParentButtons{
		individual: individual,
		document:   document,
		visibility: visibility,
	}
}

func (c *AllParentButtons) WriteTo(w io.Writer) (int64, error) {
	families := c.individual.Families()
	components := []Component{}

	for _, family := range families {
		husbandMatches := family.Husband().Is(c.individual)
		wifeMatches := family.Wife().Is(c.individual)

		if husbandMatches || wifeMatches {
			continue
		}

		components = append(components,
			NewParentButtons(c.document, family, c.visibility))
	}

	// If there are no families we still want to show an empty family. We just
	// create a dummy family that has no child nodes.
	if len(components) == 0 {
		familyNode := gedcom.NewFamilyNode(nil, "", nil)
		components = []Component{
			NewParentButtons(c.document, familyNode, c.visibility),
		}
	}

	return NewComponents(components...).WriteTo(w)
}
