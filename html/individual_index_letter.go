package html

import (
	"github.com/elliotchance/gedcom/html/core"
	"io"
	"unicode"
)

type IndividualIndexLetter struct {
	letter     rune
	isSelected bool
}

func NewIndividualIndexLetter(letter rune, isSelected bool) *IndividualIndexLetter {
	return &IndividualIndexLetter{
		letter:     letter,
		isSelected: isSelected,
	}
}

func (c *IndividualIndexLetter) WriteHTMLTo(w io.Writer) (int64, error) {
	text := string(unicode.ToUpper(c.letter))
	link := PageIndividuals(c.letter)

	return core.NewNavLink(text, link, c.isSelected).WriteHTMLTo(w)
}
