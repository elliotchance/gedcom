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

	birth := gedcom.First(c.individual.Births())
	birthDate := gedcom.First(gedcom.Dates(birth))
	birthPlace := gedcom.First(gedcom.Places(birth))
	birthDateString := gedcom.String(birthDate)
	birthPlaceString := gedcom.String(birthPlace)

	event := newIndividualEvent("Birth", birthDateString, birthPlaceString, "",
		c.document)
	events = append(events, event)

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
			event := newIndividualEvent("Marriage", date.Value(), place,
				description, c.document)
			events = append(events, event)
		}
	}

	death := gedcom.First(c.individual.Deaths())
	deathDate := gedcom.First(gedcom.Dates(death))
	deathPlace := gedcom.First(gedcom.Places(death))
	deathDateString := gedcom.String(deathDate)
	deathPlaceString := gedcom.String(deathPlace)

	individualEvent := newIndividualEvent("Death", deathDateString, deathPlaceString, "",
		c.document)
	events = append(events, individualEvent)

	tableHead := html.NewTableHead("Type", "Date", "Place", "Description")
	components := html.NewComponents(events...)
	s := html.NewTable("text-center", tableHead, components)

	return html.NewRow(html.NewColumn(html.EntireRow,
		newCard("Events", len(events), s),
	)).String()
}
