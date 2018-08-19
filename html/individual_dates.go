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
	birthDate, _ := GetBirth(c.individual)
	deathDate, _ := GetDeath(c.individual)

	eventDates := NewEventDates([]*EventDate{
		NewEventDate("b.", birthDate),
		NewEventDate("d.", deathDate),
	}).String()

	if c.individual != nil && c.individual.IsLiving() && !c.showLiving {
		eventDates = "living"
	}

	return eventDates
}
