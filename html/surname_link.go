package html

import (
	"fmt"
	"io"
	"unicode"
)

type SurnameLink struct {
	surname string
}

func NewSurnameLink(surname string) *SurnameLink {
	return &SurnameLink{
		surname: surname,
	}
}

func (c *SurnameLink) WriteTo(w io.Writer) (int64, error) {
	firstLetter := rune(c.surname[0])
	lowerFirstLetter := unicode.ToLower(firstLetter)
	destination := fmt.Sprintf("%s#%s", PageIndividuals(lowerFirstLetter), c.surname)

	return NewLink(NewText(c.surname), destination).WriteTo(w)
}
