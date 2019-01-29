package html

import (
	"github.com/elliotchance/gedcom"
	"io"
	"sort"
)

// IndividualEvents is the table of events show in the "Events" section of the
// individuals page.
type IndividualEvents struct {
	document   *gedcom.Document
	individual *gedcom.IndividualNode
	visibility LivingVisibility
}

func newIndividualEvents(document *gedcom.Document, individual *gedcom.IndividualNode, visibility LivingVisibility) *IndividualEvents {
	return &IndividualEvents{
		document:   document,
		individual: individual,
		visibility: visibility,
	}
}

func (c *IndividualEvents) WriteTo(w io.Writer) (int64, error) {
	events := []Component{}

	for _, event := range c.individual.AllEvents() {
		date := gedcom.String(gedcom.First(gedcom.Dates(event)))
		place := gedcom.String(gedcom.First(gedcom.Places(event)))

		e := NewIndividualEvent(date, place, NewEmpty(), c.individual, event)
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

		var description Component = NewEmpty()
		if family.Husband().IsIndividual(c.individual) {
			description = NewHTML(UnknownEmphasis)

			if wife := family.Wife(); wife != nil {
				description = NewIndividualLink(c.document, wife.Individual(), c.visibility)
			}
		}

		if family.Wife().IsIndividual(c.individual) {
			description = NewHTML(UnknownEmphasis)

			if husband := family.Husband(); husband != nil {
				description = NewIndividualLink(c.document, husband.Individual(), c.visibility)
			}
		}

		// Empty description means that the individual is a child so this is not
		// an event we want to show.
		if _, ok := description.(*Empty); !ok {
			event := NewIndividualEvent(date.Value(), place,
				description, c.individual, marriage)
			events = append(events, event)
		}
	}

	// Sort events by age.
	sort.Slice(events, func(i, j int) bool {
		a := events[i].(*IndividualEvent)
		b := events[j].(*IndividualEvent)

		if !a.event.Tag().Is(b.event.Tag()) {
			return a.event.Tag().SortValue() < b.event.Tag().SortValue()
		}

		aStart, _ := c.individual.AgeAt(a.event)
		bStart, _ := c.individual.AgeAt(b.event)

		return aStart.Years() < bStart.Years()
	})

	tableHead := NewTableHead("Age", "Type", "Date", "Place", "Description")
	components := NewComponents(events...)
	s := NewTable("text-center", tableHead, components)

	return NewRow(NewColumn(EntireRow,
		NewCard(NewText("Events"), len(events), s),
	)).WriteTo(w)
}
