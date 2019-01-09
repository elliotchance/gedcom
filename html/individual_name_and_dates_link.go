package html

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"io"
)

type IndividualNameAndDatesLink struct {
	individual  *gedcom.IndividualNode
	showLiving  bool
	unknownText string
}

func NewIndividualNameAndDatesLink(individual *gedcom.IndividualNode, showLiving bool, unknownText string) *IndividualNameAndDatesLink {
	return &IndividualNameAndDatesLink{
		individual:  individual,
		showLiving:  showLiving,
		unknownText: unknownText,
	}
}

func (c *IndividualNameAndDatesLink) WriteTo(w io.Writer) (int64, error) {
	if c.individual == nil {
		return writeNothing()
	}

	text := NewIndividualNameAndDates(c.individual, c.showLiving, c.unknownText)
	link := fmt.Sprintf("#%s", c.individual.Pointer())

	return NewLink(text, link).Style("color: black").WriteTo(w)
}
