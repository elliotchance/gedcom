package html

import (
	"github.com/elliotchance/gedcom"
	"io"
)

type IndividualNameAndDates struct {
	individual  *gedcom.IndividualNode
	showLiving  bool
	unknownText string
}

func NewIndividualNameAndDates(individual *gedcom.IndividualNode, showLiving bool, unknownText string) *IndividualNameAndDates {
	return &IndividualNameAndDates{
		individual:  individual,
		showLiving:  showLiving,
		unknownText: unknownText,
	}
}

func (c *IndividualNameAndDates) WriteTo(w io.Writer) (int64, error) {
	name := NewIndividualName(c.individual, c.showLiving, c.unknownText)
	dates := NewIndividualDates(c.individual, c.showLiving)

	if name.IsUnknown() || dates.IsBlank() {
		return name.WriteTo(w)
	}

	return NewComponents(name, NewText(" ("), dates, NewText(")")).WriteTo(w)
}
