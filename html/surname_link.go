package html

import (
	"fmt"
	"io"
	"unicode"

	"github.com/elliotchance/gedcom/v39/html/core"
)

type SurnameLink struct {
	surname string
}

func NewSurnameLink(surname string) *SurnameLink {
	return &SurnameLink{
		surname: surname,
	}
}

func (c *SurnameLink) WriteHTMLTo(w io.Writer) (int64, error) {
	firstLetter := rune(c.surname[0])
	lowerFirstLetter := unicode.ToLower(firstLetter)
	destination := fmt.Sprintf("%s#%s", PageIndividuals(lowerFirstLetter), c.surname)

	return core.NewLink(core.NewText(c.surname), destination).WriteHTMLTo(w)
}
