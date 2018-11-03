package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
	"sort"
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

	for _, event := range c.individual.AllEvents() {
		date := gedcom.String(gedcom.First(gedcom.Dates(event)))
		place := gedcom.String(gedcom.First(gedcom.Places(event)))

		e := newIndividualEvent(date, place, html.NewEmpty(), c.individual, event)
		events = append(events, e)
	}

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

		var description fmt.Stringer = html.NewEmpty()
		if family.Husband().Is(c.individual) {
			description = html.NewHTML("<em>Unknown</em>")

			if wife := family.Wife(); wife != nil {
				description = newIndividualLink(c.document, wife)
			}
		}

		if family.Wife().Is(c.individual) {
			description = html.NewHTML("<em>Unknown</em>")

			if husband := family.Husband(); husband != nil {
				description = newIndividualLink(c.document, husband)
			}
		}

		// Empty description means that the individual is a child so this is not
		// an event we want to show.
		if _, ok := description.(*html.Empty); !ok {
			event := newIndividualEvent(date.Value(), place,
				description, c.individual, marriage)
			events = append(events, event)
		}
	}

	// Sort events by age.
	sort.Slice(events, func(i, j int) bool {
		a := events[i].(*individualEvent)
		b := events[j].(*individualEvent)

		if !a.event.Tag().Is(b.event.Tag()) {
			return a.event.Tag().SortValue() < b.event.Tag().SortValue()
		}

		aStart, _ := c.individual.AgeAt(a.event)
		bStart, _ := c.individual.AgeAt(b.event)

		return aStart.Years() < bStart.Years()
	})

	tableHead := html.NewTableHead("Age", "Type", "Date", "Place", "Description")
	components := html.NewComponents(events...)
	s := html.NewTable("text-center", tableHead, components)

	return html.NewRow(html.NewColumn(html.EntireRow,
		newCard("Events", len(events), s),
	)).String()
}
