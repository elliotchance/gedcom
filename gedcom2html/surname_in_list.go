package main

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
)

type surnameInList struct {
	document *gedcom.Document
	surname  string
}

func newSurnameInList(document *gedcom.Document, surname string) *surnameInList {
	return &surnameInList{
		document: document,
		surname:  surname,
	}
}

func (c *surnameInList) String() string {
	count := 0
	for _, individual := range c.document.Individuals() {
		if individual.Name().Surname() == c.surname {
			count++
		}
	}

	return html.NewTableRow(
		html.NewTableCell(newSurnameLink(c.surname)),
		html.NewTableCell(html.NewNumber(count)),
	).String()
}
