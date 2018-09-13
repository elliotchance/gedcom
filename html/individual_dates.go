package html

import (
	"github.com/elliotchance/gedcom"
)

type IndividualDates struct {
	individual *gedcom.IndividualNode
	showLiving bool
}

func NewIndividualDates(individual *gedcom.IndividualNode, showLiving bool) *IndividualDates {
	return &IndividualDates{
		individual: individual,
		showLiving: showLiving,
	}
}

func (c *IndividualDates) String() string {
	birthDate := gedcom.First(c.individual.Births())
	deathDate := gedcom.Last(c.individual.Deaths())

	eventDates := NewEventDates([]*EventDate{
		NewEventDate("b.", gedcom.String(birthDate)),
		NewEventDate("d.", gedcom.String(deathDate)),
	}).String()

	if c.individual != nil && c.individual.IsLiving() && !c.showLiving {
		eventDates = "living"
	}

	return eventDates
}
