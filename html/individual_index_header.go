package html

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
	"io"
	"strings"
)

type IndividualIndexHeader struct {
	document         *gedcom.Document
	selectedLetter   rune
	livingVisibility LivingVisibility
	indexLetters     []rune
}

func NewIndividualIndexHeader(document *gedcom.Document, selectedLetter rune, livingVisibility LivingVisibility, indexLetters []rune) *IndividualIndexHeader {
	return &IndividualIndexHeader{
		document:         document,
		selectedLetter:   selectedLetter,
		livingVisibility: livingVisibility,
		indexLetters:     indexLetters,
	}
}

func GetIndexLetters(document *gedcom.Document, livingVisibility LivingVisibility) []rune {
	letterMap := map[rune]bool{}
	for _, individual := range document.Individuals() {
		switch livingVisibility {
		case LivingVisibilityShow, LivingVisibilityPlaceholder:
			letterMap[getIndexLetter(individual)] = true
		case LivingVisibilityHide:
			// nothing
		}
	}

	indexLetters := []rune{}
	if _, ok := letterMap[symbolLetter]; ok {
		indexLetters = []rune{symbolLetter}
	}

	for i := rune('a'); i <= rune('z'); i++ {
		if _, ok := letterMap[i]; ok {
			indexLetters = append(indexLetters, i)
		}
	}

	return indexLetters
}

func getIndexLetter(individual *gedcom.IndividualNode) rune {
	name := strings.ToLower(individual.Name().String())

	switch {
	case name == "", name[0] < 'a', name[0] > 'z':
		return symbolLetter
	}

	return rune(name[0])
}

func (c *IndividualIndexHeader) WriteHTMLTo(w io.Writer) (int64, error) {
	pills := []core.Component{}

	for _, letter := range c.indexLetters {
		pills = append(pills,
			NewIndividualIndexLetter(letter, letter == c.selectedLetter))
	}

	return core.NewRow(
		core.NewColumn(core.EntireRow, core.NewNavPills(pills)),
	).WriteHTMLTo(w)
}
