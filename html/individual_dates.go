package html

import (
	"github.com/elliotchance/gedcom"
	"io"
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

func (c *IndividualDates) IsBlank() bool {
	eventDates := c.EventDates()

	for _, eventDate := range eventDates {
		if !eventDate.IsBlank() {
			return false
		}
	}

	return true
}

func (c *IndividualDates) WriteTo(w io.Writer) (int64, error) {
	if c.individual != nil && c.individual.IsLiving() && !c.showLiving {
		return NewText("living").WriteTo(w)
	}

	eventDates := c.EventDates()

	return NewEventDates(eventDates).WriteTo(w)
}

func (c *IndividualDates) EventDates() []*EventDate {
	eventDates := []*EventDate{}

	// Use birth or fallback to baptism.
	births := c.individual.Births()
	baptisms := c.individual.Baptisms()
	switch {
	case len(births) > 0:
		eventDate := NewEventDate("b.", births[0].Dates())
		eventDates = append(eventDates, eventDate)

	case len(baptisms) > 0:
		eventDate := NewEventDate("bap.", baptisms[0].Dates())
		eventDates = append(eventDates, eventDate)
	}

	// Use death or fallback to burial.
	deaths := c.individual.Deaths()
	burials := c.individual.Burials()
	switch {
	case len(deaths) > 0:
		eventDate := NewEventDate("d.", deaths[0].Dates())
		eventDates = append(eventDates, eventDate)

	case len(burials) > 0:
		eventDate := NewEventDate("bur.", burials[0].Dates())
		eventDates = append(eventDates, eventDate)
	}

	return eventDates
}
