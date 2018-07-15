package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
)

// individualButton is a large coloured button that links to an individuals
// page. It contains the same and some date information. This is also used to
// represent unknown or missing individuals.
type individualButton struct {
	individual *gedcom.IndividualNode
}

func newIndividualButton(individual *gedcom.IndividualNode) *individualButton {
	return &individualButton{
		individual: individual,
	}
}

func (c *individualButton) String() string {
	name := newIndividualName(c.individual).String()
	birthDate, _ := getBirth(c.individual)
	deathDate, _ := getDeath(c.individual)

	onclick := ""
	if c.individual != nil {
		onclick = fmt.Sprintf(`onclick="location.href='%s'"`,
			pageIndividual(c.individual))
	}

	eventDates := newEventDates([]*eventDate{
		newEventDate("b.", birthDate),
		newEventDate("d.", deathDate),
	}).String()

	// If the individual is living we need to hide all their information.
	if c.individual != nil && c.individual.IsLiving() {
		name = "<em>Hidden</em>"
		onclick = ""
		eventDates = "living"
	}

	return fmt.Sprintf(`
		<button type="button" class="btn btn-outline-%s btn-block" %s>
			<strong>%s</strong><br/>
			%s&nbsp;
		</button>
	`, colorClassForIndividual(c.individual), onclick, name, eventDates)
}
