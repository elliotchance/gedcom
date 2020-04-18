package html

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
	"io"
	"sort"
)

// IndividualEvents is the table of events show in the "Events" section of the
// individuals page.
type IndividualEvents struct {
	document   *gedcom.Document
	individual *gedcom.IndividualNode
	visibility LivingVisibility
	placesMap  map[string]*place
}

func NewIndividualEvents(document *gedcom.Document, individual *gedcom.IndividualNode, visibility LivingVisibility, placesMap map[string]*place) *IndividualEvents {
	return &IndividualEvents{
		document:   document,
		individual: individual,
		visibility: visibility,
		placesMap:  placesMap,
	}
}

func (c *IndividualEvents) WriteHTMLTo(w io.Writer) (int64, error) {
	var events []core.Component

	for _, event := range c.individual.AllEvents() {
		date, place := gedcom.DateAndPlace(event)
		//date := gedcom.String(gedcom.First(gedcom.Dates(event)))
		//place := gedcom.String(gedcom.First(gedcom.Places(event)))

		e := NewIndividualEvent(gedcom.String(date), gedcom.String(place),
			core.NewEmpty(), c.individual, event, c.placesMap)
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

		var description core.Component = core.NewEmpty()
		if family.Husband().IsIndividual(c.individual) {
			description = core.NewHTML(UnknownEmphasis)

			if wife := family.Wife(); wife != nil {
				description = NewIndividualLink(c.document, wife.Individual(),
					c.visibility, c.placesMap)
			}
		}

		if family.Wife().IsIndividual(c.individual) {
			description = core.NewHTML(UnknownEmphasis)

			if husband := family.Husband(); husband != nil {
				description = NewIndividualLink(c.document,
					husband.Individual(), c.visibility, c.placesMap)
			}
		}

		// Empty description means that the individual is a child so this is not
		// an event we want to show.
		if _, ok := description.(*core.Empty); !ok {
			event := NewIndividualEvent(date.Value(), place,
				description, c.individual, marriage, c.placesMap)
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

	tableHead := core.NewTableHead("Age", "Type", "Date", "Place", "Description")
	components := core.NewComponents(events...)
	s := core.NewTable("text-center", tableHead, components)

	return core.NewRow(core.NewColumn(core.EntireRow,
		core.NewCard(core.NewText("Events"), len(events), s),
	)).WriteHTMLTo(w)
}
