package html

import (
	"github.com/elliotchance/gedcom"
)

const UnknownEmphasis = "<em>Unknown</em>"

// IndividualName outputs the full name of the individual. This is a wrapper for
// the String function on the IndividualNode. If the individual does not have
// any names then "Unknown" will be used. It is safe to use nil for the
// individual.
type IndividualName struct {
	individual  *gedcom.IndividualNode
	showLiving  bool
	unknownText string
}

func NewIndividualName(individual *gedcom.IndividualNode, showLiving bool, unknownText string) *IndividualName {
	return &IndividualName{
		individual:  individual,
		showLiving:  showLiving,
		unknownText: unknownText,
	}
}

func (c *IndividualName) String() string {
	if c.individual == nil {
		return c.unknownText
	}

	if c.individual.IsLiving() && !c.showLiving {
		return "<em>Hidden</em>"
	}

	names := c.individual.Names()
	if len(names) == 0 {
		return c.unknownText
	}

	return names[0].String()
}
