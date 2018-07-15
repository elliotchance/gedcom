package main

import "fmt"

// eventDate shows a date like "d. 1882" but will not show anything if the date
// is not provided.
type eventDate struct {
	event string
	date  string
}

func newEventDate(event, date string) *eventDate {
	return &eventDate{
		event: event,
		date:  date,
	}
}

func (c *eventDate) String() string {
	if c.date == "" {
		return ""
	}

	return fmt.Sprintf("<em>%s</em> %s", c.event, c.date)
}
