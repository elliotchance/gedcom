package html

import (
	"github.com/elliotchance/gedcom"
	"io"
)

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

func (c *EventDate) WriteTo(w io.Writer) (int64, error) {
	if c.IsBlank() {
		return writeNothing()
	}

	return NewComponents(
		NewTag("em", nil, NewText(c.event)),
		NewText(" "+c.dates[0].String()),
	).WriteTo(w)
}

func (c *EventDate) IsBlank() bool {
	return len(c.dates) == 0
}
