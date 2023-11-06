package html

import (
	"io"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/gedcom/v39/html/core"
)

const UnknownEmphasis = "<em>Unknown</em>"

// IndividualName outputs the full name of the individual. This is a wrapper for
// the String function on the IndividualNode. If the individual does not have
// any names then "Unknown" will be used. It is safe to use nil for the
// individual.
type IndividualName struct {
	individual  *gedcom.IndividualNode
	visibility  LivingVisibility
	unknownHTML string
}

func NewIndividualName(individual *gedcom.IndividualNode, visibility LivingVisibility, unknownHTML string) *IndividualName {
	return &IndividualName{
		individual:  individual,
		visibility:  visibility,
		unknownHTML: unknownHTML,
	}
}

func (c *IndividualName) IsUnknown() bool {
	if c.individual == nil {
		return true
	}

	names := c.individual.Names()

	return len(names) == 0
}

func (c *IndividualName) WriteHTMLTo(w io.Writer) (int64, error) {
	if c.individual == nil {
		return writeString(w, c.unknownHTML)
	}

	isLiving := c.individual.IsLiving()
	if isLiving {
		switch c.visibility {
		case LivingVisibilityShow:
			// Proceed.

		case LivingVisibilityHide:
			return writeNothing()

		case LivingVisibilityPlaceholder:
			return writeString(w, "<em>Hidden</em>")
		}
	}

	names := c.individual.Names()
	if len(names) == 0 {
		return writeString(w, c.unknownHTML)
	}

	return core.NewText(names[0].String()).WriteHTMLTo(w)
}
