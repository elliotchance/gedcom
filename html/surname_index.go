package html

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/util"
	"io"
	"sort"
)

type SurnameIndex struct {
	document       *gedcom.Document
	selectedLetter rune
}

func NewSurnameIndex(document *gedcom.Document, selectedLetter rune) *SurnameIndex {
	return &SurnameIndex{
		document:       document,
		selectedLetter: selectedLetter,
	}
}

func (c *SurnameIndex) WriteTo(w io.Writer) (int64, error) {
	surnames := []string{}

	for _, individual := range c.document.Individuals() {
		if individual.IsLiving() {
			continue
		}

		surname := individual.Name().Surname()
		exists := util.StringSliceContains(surnames, surname)
		if surnameStartsWith(individual, c.selectedLetter) && !exists {
			surnames = append(surnames, surname)
		}
	}

	// Sort surnames
	sort.Strings(surnames)

	// Render
	pills := []Component{}
	for _, surname := range surnames {
		pills = append(pills, NewNavLink(surname, "#"+surname, false))
	}

	return NewNavPillsRow(pills).WriteTo(w)
}
