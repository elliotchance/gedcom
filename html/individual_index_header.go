package html

import (
	"github.com/elliotchance/gedcom"
	"io"
	"strings"
)

type IndividualIndexHeader struct {
	document       *gedcom.Document
	selectedLetter rune
}

func NewIndividualIndexHeader(document *gedcom.Document, selectedLetter rune) *IndividualIndexHeader {
	return &IndividualIndexHeader{
		document:       document,
		selectedLetter: selectedLetter,
	}
}

func GetIndexLetters(document *gedcom.Document) []rune {
	letterMap := map[rune]bool{}
	for _, individual := range document.Individuals() {
		letterMap[getIndexLetter(individual)] = true
	}

	letters := []rune{}
	if _, ok := letterMap[symbolLetter]; ok {
		letters = []rune{symbolLetter}
	}

	for i := rune('a'); i <= rune('z'); i++ {
		if _, ok := letterMap[i]; ok {
			letters = append(letters, i)
		}
	}

	return letters
}

func getIndexLetter(individual *gedcom.IndividualNode) rune {
	name := strings.ToLower(individual.Name().String())

	switch {
	case name == "", name[0] < 'a', name[0] > 'z':
		name = "#"
	}

	return rune(name[0])
}

func (c *IndividualIndexHeader) WriteTo(w io.Writer) (int64, error) {
	pills := []Component{}

	for _, letter := range GetIndexLetters(c.document) {
		pills = append(pills,
			NewIndividualIndexLetter(letter, letter == c.selectedLetter))
	}

	return NewRow(
		NewColumn(EntireRow, NewNavPills(pills)),
	).WriteTo(w)
}
