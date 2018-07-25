package main

import (
	"github.com/elliotchance/gedcom"
)

// individualName outputs the full name of the individual. This is a wrapper for
// the String function on the IndividualNode. If the individual does not have
// any names then "Unknown" will be used. It is safe to use nil for the
// individual.
type individualName struct {
	individual *gedcom.IndividualNode
}

func newIndividualName(individual *gedcom.IndividualNode) *individualName {
	return &individualName{
		individual: individual,
	}
}

func (c *individualName) String() string {
	if c.individual == nil {
		return "<em>Unknown</em>"
	}

	if c.individual.IsLiving() {
		return "<em>Hidden</em>"
	}

	names := c.individual.Names()
	if len(names) == 0 {
		return "<em>Unknown</em>"
	}

	return names[0].String()
}
