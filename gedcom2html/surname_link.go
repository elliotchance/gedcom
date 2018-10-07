package main

import (
	"fmt"
	"github.com/elliotchance/gedcom/html"
	"unicode"
)

type surnameLink struct {
	surname string
}

func newSurnameLink(surname string) *surnameLink {
	return &surnameLink{
		surname: surname,
	}
}

func (c *surnameLink) String() string {
	firstLetter := rune(c.surname[0])
	lowerFirstLetter := unicode.ToLower(firstLetter)
	destination := fmt.Sprintf("%s#%s", pageIndividuals(lowerFirstLetter), c.surname)

	return html.NewLink(c.surname, destination).String()
}
