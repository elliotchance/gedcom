package html

import (
	"io"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/gedcom/v39/html/core"
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

func (c *SurnameInList) WriteHTMLTo(w io.Writer) (int64, error) {
	count := 0
	for _, individual := range c.document.Individuals() {
		if individual.Name().Surname() == c.surname {
			count++
		}
	}

	return core.NewTableRow(
		core.NewTableCell(NewSurnameLink(c.surname)),
		core.NewTableCell(core.NewNumber(count)),
	).WriteHTMLTo(w)
}
