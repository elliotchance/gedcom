package main

import "fmt"

// individualEvent is a row in the "Events" section of the individuals page.
type individualEvent struct {
	kind        string
	date        string
	place       string
	description string
}

func newIndividualEvent(kind, date, place, description string) *individualEvent {
	return &individualEvent{
		kind:        kind,
		date:        date,
		place:       place,
		description: description,
	}
}

func (c *individualEvent) String() string {
	return fmt.Sprintf(`
		<tr>
			<th>%s</th>
			<td>%s</td>
			<td>%s</td>
			<td>%s</td>
		</tr>`, c.kind, c.date, prettyPlaceName(c.place), c.description)
}
