package main

import (
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
	return html.Sprintf(`
		<a href="%s#%s">%s</a>`, pageIndividuals(unicode.ToLower(rune(c.surname[0]))), c.surname, c.surname)
}
