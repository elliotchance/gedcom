package main

// eventDates contains several eventDate instances that are separated by three
// spaces. Only the dates that are non-empty will be shown.
type eventDates struct {
	items []*eventDate
}

func newEventDates(items []*eventDate) *eventDates {
	return &eventDates{
		items: items,
	}
}

func (c *eventDates) String() string {
	s := ""
	for _, date := range c.items {
		if s != "" && date.String() != "" {
			s += "&nbsp;&nbsp;&nbsp;"
		}

		s += date.String()
	}

	return s
}
