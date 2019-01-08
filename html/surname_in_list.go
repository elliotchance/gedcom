package html

import (
	"github.com/elliotchance/gedcom"
	"io"
)

type SurnameInList struct {
	document *gedcom.Document
	surname  string
}

func NewSurnameInList(document *gedcom.Document, surname string) *SurnameInList {
	return &SurnameInList{
		document: document,
		surname:  surname,
	}
}

func (c *SurnameInList) WriteTo(w io.Writer) (int64, error) {
	count := 0
	for _, individual := range c.document.Individuals() {
		if individual.Name().Surname() == c.surname {
			count++
		}
	}

	return NewTableRow(
		NewTableCell(NewSurnameLink(c.surname)),
		NewTableCell(NewNumber(count)),
	).WriteTo(w)
}
