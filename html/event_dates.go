package html

import "io"

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

func (c *EventDates) WriteHTMLTo(w io.Writer) (int64, error) {
	n := int64(0)
	for _, date := range c.items {
		if n > 0 && !date.IsBlank() {
			n += appendString(w, "&nbsp;&nbsp;&nbsp;")
		}

		n += appendComponent(w, date)
	}

	return n, nil
}
