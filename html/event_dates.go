package html

// EventDates contains several eventDate instances that are separated by three
// spaces. Only the dates that are non-empty will be shown.
type EventDates struct {
	items []*EventDate
}

func NewEventDates(items []*EventDate) *EventDates {
	return &EventDates{
		items: items,
	}
}

func (c *EventDates) String() string {
	s := ""
	for _, date := range c.items {
		if s != "" && date.String() != "" {
			s += "&nbsp;&nbsp;&nbsp;"
		}

		s += date.String()
	}

	return s
}
