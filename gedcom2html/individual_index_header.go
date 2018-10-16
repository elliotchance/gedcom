package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
	"strings"
)

type individualIndexHeader struct {
	document       *gedcom.Document
	selectedLetter rune
}

func newIndividualIndexHeader(document *gedcom.Document, selectedLetter rune) *individualIndexHeader {
	return &individualIndexHeader{
		document:       document,
		selectedLetter: selectedLetter,
	}
}

func getIndexLetters(document *gedcom.Document) []rune {
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

func (c *individualIndexHeader) String() string {
	pills := []fmt.Stringer{}

	for _, letter := range getIndexLetters(c.document) {
		pills = append(pills, newIndividualIndexLetter(letter, letter == c.selectedLetter))
	}

	return html.NewRow(
		html.NewColumn(html.EntireRow, newNavPills(pills)),
	).String()
}
