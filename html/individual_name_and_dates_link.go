package html

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"io"
)

type IndividualNameAndDatesLink struct {
	individual  *gedcom.IndividualNode
	visibility  LivingVisibility
	unknownText string
}

func NewIndividualNameAndDatesLink(individual *gedcom.IndividualNode, visibility LivingVisibility, unknownText string) *IndividualNameAndDatesLink {
	return &IndividualNameAndDatesLink{
		individual:  individual,
		visibility:  visibility,
		unknownText: unknownText,
	}
}

func (c *IndividualNameAndDatesLink) WriteTo(w io.Writer) (int64, error) {
	if c.individual == nil {
		return writeNothing()
	}

	text := NewIndividualNameAndDates(c.individual, c.visibility, c.unknownText)
	link := fmt.Sprintf("#%s", c.individual.Pointer())

	return NewLink(text, link).Style("color: black").WriteTo(w)
}
