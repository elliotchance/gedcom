package html

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
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

func (c *EventDate) WriteHTMLTo(w io.Writer) (int64, error) {
	if c.IsBlank() {
		return writeNothing()
	}

	return core.NewComponents(
		core.NewTag("em", nil, core.NewText(c.event)),
		core.NewText(" "+c.dates[0].String()),
	).WriteHTMLTo(w)
}

func (c *EventDate) IsBlank() bool {
	return len(c.dates) == 0
}
