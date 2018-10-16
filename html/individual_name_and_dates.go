package html

import (
	"github.com/elliotchance/gedcom"
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

func (c *IndividualNameAndDates) String() string {
	name := NewIndividualName(c.individual, c.showLiving, c.unknownText).String()
	dates := NewIndividualDates(c.individual, c.showLiving).String()

	if name == c.unknownText || dates == "" {
		return name
	}

	return Sprintf("%s (%s)", name, dates)
}
