package html

import (
	"github.com/elliotchance/gedcom"
	"io"
)

const UnknownEmphasis = "<em>Unknown</em>"

// IndividualName outputs the full name of the individual. This is a wrapper for
// the String function on the IndividualNode. If the individual does not have
// any names then "Unknown" will be used. It is safe to use nil for the
// individual.
type IndividualName struct {
	individual  *gedcom.IndividualNode
	showLiving  bool
	unknownHTML string
}

func NewIndividualName(individual *gedcom.IndividualNode, showLiving bool, unknownHTML string) *IndividualName {
	return &IndividualName{
		individual:  individual,
		showLiving:  showLiving,
		unknownHTML: unknownHTML,
	}
}

func (c *IndividualName) IsUnknown() bool {
	return c.individual == nil || len(c.individual.Names()) == 0
}

func (c *IndividualName) WriteTo(w io.Writer) (int64, error) {
	if c.individual == nil {
		return writeString(w, c.unknownHTML)
	}

	isLiving := c.individual.IsLiving()
	if isLiving && !c.showLiving {
		return writeString(w, "<em>Hidden</em>")
	}

	names := c.individual.Names()
	if len(names) == 0 {
		return writeString(w, c.unknownHTML)
	}

	return NewText(names[0].String()).WriteTo(w)
}
