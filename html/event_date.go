package html

import "fmt"

// EventDate shows a date like "d. 1882" but will not show anything if the date
// is not provided.
type EventDate struct {
	event string
	date  string
}

func NewEventDate(event, date string) *EventDate {
	return &EventDate{
		event: event,
		date:  date,
	}
}

func (c *EventDate) String() string {
	if c.date == "" {
		return ""
	}

	return fmt.Sprintf("<em>%s</em> %s", c.event, c.date)
}
