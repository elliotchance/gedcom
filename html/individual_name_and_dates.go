package html

import (
	"github.com/elliotchance/gedcom"
	"io"
)

type IndividualNameAndDates struct {
	individual  *gedcom.IndividualNode
	visibility  LivingVisibility
	unknownText string
}

func NewIndividualNameAndDates(individual *gedcom.IndividualNode, visibility LivingVisibility, unknownText string) *IndividualNameAndDates {
	return &IndividualNameAndDates{
		individual:  individual,
		visibility:  visibility,
		unknownText: unknownText,
	}
}

func (c *IndividualNameAndDates) WriteTo(w io.Writer) (int64, error) {
	name := NewIndividualName(c.individual, c.visibility, c.unknownText)
	dates := NewIndividualDates(c.individual, c.visibility)

	if name.IsUnknown() || dates.IsBlank() {
		return name.WriteTo(w)
	}

	return NewComponents(name, NewText(" ("), dates, NewText(")")).WriteTo(w)
}
