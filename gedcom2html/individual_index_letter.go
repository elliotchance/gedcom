package main

import "unicode"

type individualIndexLetter struct {
	letter     rune
	isSelected bool
}

func newIndividualIndexLetter(letter rune, isSelected bool) *individualIndexLetter {
	return &individualIndexLetter{
		letter:     letter,
		isSelected: isSelected,
	}
}

func (c *individualIndexLetter) String() string {
	text := string(unicode.ToUpper(c.letter))
	link := pageIndividuals(c.letter)

	return newNavLink(text, link, c.isSelected).String()
}
