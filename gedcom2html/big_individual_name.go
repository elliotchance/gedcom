package main

import (
	"github.com/elliotchance/gedcom"
)

// bigIndividualName shows the large individual name on their individual page
// (right below their parents).
type bigIndividualName struct {
	individual *gedcom.IndividualNode
}

func newBigIndividualName(individual *gedcom.IndividualNode) *bigIndividualName {
	return &bigIndividualName{
		individual: individual,
	}
}

func (c *bigIndividualName) String() string {
	name := newIndividualName(c.individual).String()

	return newRow(
		newColumn(entireRow,
			newHeading(1, "text-center", name),
		),
	).String()
}
