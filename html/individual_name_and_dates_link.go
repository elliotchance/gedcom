package html

import (
	"github.com/elliotchance/gedcom"
	"fmt"
)

type IndividualNameAndDatesLink struct {
	individual  *gedcom.IndividualNode
	showLiving  bool
	unknownText string
}

func NewIndividualNameAndDatesLink(individual *gedcom.IndividualNode, showLiving bool, unknownText string) *IndividualNameAndDatesLink {
	return &IndividualNameAndDatesLink{
		individual:  individual,
		showLiving:  showLiving,
		unknownText: unknownText,
	}
}

func (c *IndividualNameAndDatesLink) String() string {
	if c.individual == nil {
		return ""
	}

	text := NewIndividualNameAndDates(c.individual, c.showLiving, c.unknownText).String()
	link := fmt.Sprintf("#%s", c.individual.Pointer())

	return NewLink(text, link).Style("color: black").String()
}
