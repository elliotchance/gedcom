package main

import (
	"github.com/elliotchance/gedcom"
)

// allParentButtons represent one or more families that an individual belongs
// to. These are show as large buttons above the large name of the person in on
// their individual page.
type allParentButtons struct {
	individual *gedcom.IndividualNode
	document   *gedcom.Document
}

func newAllParentButtons(document *gedcom.Document, individual *gedcom.IndividualNode) *allParentButtons {
	return &allParentButtons{
		individual: individual,
		document:   document,
	}
}

func (c *allParentButtons) String() (s string) {
	families := c.individual.Families()

	for _, family := range families {
		husbandMatches := family.Husband().Is(c.individual)
		wifeMatches := family.Wife().Is(c.individual)

		if husbandMatches || wifeMatches {
			continue
		}

		s += newParentButtons(c.document, family).String()
	}

	// If there are no families we still want to show an empty family. We just
	// create a dummy family that has no child nodes.
	if s == "" {
		familyNode := gedcom.NewFamilyNode(nil, "", nil)
		s = newParentButtons(c.document, familyNode).String()
	}

	return
}
