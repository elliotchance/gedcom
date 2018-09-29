package html

import "github.com/elliotchance/gedcom"

// EventDate shows a date like "d. 1882" but will not show anything if the date
// is not provided.
type EventDate struct {
	event string
	dates []*gedcom.DateNode
}

func NewEventDate(event string, dates []*gedcom.DateNode) *EventDate {
	return &EventDate{
		event: event,
		dates: dates,
	}
}

func (c *EventDate) String() string {
	if len(c.dates) == 0 {
		return ""
	}

	return Sprintf("<em>%s</em> %s", c.event, c.dates[0].String())
}
