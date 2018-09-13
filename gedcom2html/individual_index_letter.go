package main

import (
	"github.com/elliotchance/gedcom/html"
	"unicode"
)

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
	active := ""
	if c.isSelected {
		active = "active"
	}

	return html.Sprintf(`
			<li class="nav-item">
    			<a class="nav-link %s" href="%s">%c</a>
  			</li>`,
		active, pageIndividuals(c.letter), unicode.ToUpper(c.letter))
}
