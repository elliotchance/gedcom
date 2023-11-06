package html

import (
	"io"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/gedcom/v39/html/core"
)

type IndividualNameAndDates struct {
	individual  *gedcom.IndividualNode
	visibility  LivingVisibility
	unknownText string
}

func NewIndividualNameAndDates(individual *gedcom.IndividualNode, visibility LivingVisibility, unknownText string) *IndividualNameAndDates {
	return &IndividualNameAndDates{
		individual:  individual,
		visibility:  visibility,
		unknownText: unknownText,
	}
}

func (c *IndividualNameAndDates) WriteHTMLTo(w io.Writer) (int64, error) {
	name := NewIndividualName(c.individual, c.visibility, c.unknownText)
	dates := NewIndividualDates(c.individual, c.visibility)

	isUnknown := name.IsUnknown()
	datesAreBlank := dates.IsBlank()

	if isUnknown || datesAreBlank {
		return name.WriteHTMLTo(w)
	}

	return core.NewComponents(name, core.NewText(" ("), dates,
		core.NewText(")")).WriteHTMLTo(w)
}
