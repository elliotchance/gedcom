package html

import (
	"fmt"
	"io"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/gedcom/v39/html/core"
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

func (c *IndividualNameAndDatesLink) WriteHTMLTo(w io.Writer) (int64, error) {
	if c.individual == nil {
		return writeNothing()
	}

	text := NewIndividualNameAndDates(c.individual, c.visibility, c.unknownText)
	link := fmt.Sprintf("#%s", c.individual.Pointer())

	return core.NewLink(text, link).Style("color: black").WriteHTMLTo(w)
}
