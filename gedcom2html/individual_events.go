package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
)

// individualEvents is the table of events show in the "Events" section of the
// individuals page.
type individualEvents struct {
	document   *gedcom.Document
	individual *gedcom.IndividualNode
}

func newIndividualEvents(document *gedcom.Document, individual *gedcom.IndividualNode) *individualEvents {
	return &individualEvents{
		document:   document,
		individual: individual,
	}
}

func (c *individualEvents) String() string {
	events := []fmt.Stringer{}

	birthDate, birthPlace := getBirth(c.individual)
	events = append(events, newIndividualEvent("Birth", birthDate, birthPlace, ""))

	for _, family := range c.individual.Families(c.document) {
		marriage := family.FirstNodeWithTag(gedcom.TagMarriage)
		if marriage == nil {
			continue
		}

		date := marriage.FirstNodeWithTag(gedcom.TagDate)
		if date == nil {
			continue
		}

		place := ""
		if p := marriage.FirstNodeWithTag(gedcom.TagPlace); p != nil {
			place = p.Value()
		}

		description := ""
		if family.Husband(c.document).Is(c.individual) {
			description = "<em>Unknown</em>"

			if wife := family.Wife(c.document); wife != nil {
				description = newIndividualLink(c.document, wife).String()
			}
		}

		if family.Wife(c.document).Is(c.individual) {
			description = "<em>Unknown</em>"

			if husband := family.Husband(c.document); husband != nil {
				description = newIndividualLink(c.document, husband).String()
			}
		}

		// Empty description means that the individual is a child so this is not
		// an event we want to show.
		if description != "" {
			events = append(events,
				newIndividualEvent("Marriage", date.Value(), place, description))
		}
	}

	deathDate, deathPlace := getDeath(c.individual)
	events = append(events, newIndividualEvent("Death", deathDate, deathPlace, ""))

	s := newTable("text-center",
		newTableHead("Type", "Date", "Place", "Description"),
		newComponents(events...))

	return newRow(newColumn(entireRow,
		newCard("Events", len(events), s),
	)).String()
}
