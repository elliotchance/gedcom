package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
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

	birthDate, birthPlace := html.GetBirth(c.individual)
	events = append(events, newIndividualEvent("Birth", birthDate, birthPlace, ""))

	for _, family := range c.individual.Families() {
		marriage := gedcom.First(gedcom.NodesWithTag(family, gedcom.TagMarriage))
		if marriage == nil {
			continue
		}

		date := gedcom.First(gedcom.NodesWithTag(marriage, gedcom.TagDate))
		if date == nil {
			continue
		}

		place := ""
		if p := gedcom.First(gedcom.NodesWithTag(marriage, gedcom.TagPlace)); p != nil {
			place = p.Value()
		}

		description := ""
		if family.Husband().Is(c.individual) {
			description = "<em>Unknown</em>"

			if wife := family.Wife(); wife != nil {
				description = newIndividualLink(c.document, wife).String()
			}
		}

		if family.Wife().Is(c.individual) {
			description = "<em>Unknown</em>"

			if husband := family.Husband(); husband != nil {
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

	deathDate, deathPlace := html.GetDeath(c.individual)
	events = append(events, newIndividualEvent("Death", deathDate, deathPlace, ""))

	s := html.NewTable("text-center",
		html.NewTableHead("Type", "Date", "Place", "Description"),
		html.NewComponents(events...))

	return html.NewRow(html.NewColumn(html.EntireRow,
		newCard("Events", len(events), s),
	)).String()
}
