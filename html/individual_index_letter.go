package html

import (
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

func (c *IndividualIndexLetter) WriteTo(w io.Writer) (int64, error) {
	text := string(unicode.ToUpper(c.letter))
	link := PageIndividuals(c.letter)

	return NewNavLink(text, link, c.isSelected).WriteTo(w)
}
