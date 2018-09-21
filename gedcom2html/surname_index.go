package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/util"
	"sort"
)

type surnameIndex struct {
	document       *gedcom.Document
	selectedLetter rune
}

func newSurnameIndex(document *gedcom.Document, selectedLetter rune) *surnameIndex {
	return &surnameIndex{
		document:       document,
		selectedLetter: selectedLetter,
	}
}

func (c *surnameIndex) String() string {
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
	pills := []fmt.Stringer{}
	for _, surname := range surnames {
		pills = append(pills, newNavLink(surname, "#"+surname, false))
	}

	return newNavPillsRow(pills).String()
}
