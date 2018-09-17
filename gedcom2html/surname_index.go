package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
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

func stringInSlice(slice []string, s string) bool {
	for _, e := range slice {
		if s == e {
			return true
		}
	}

	return false
}

func (c *surnameIndex) String() string {
	surnames := []string{}

	for _, individual := range c.document.Individuals() {
		if individual.IsLiving() {
			continue
		}

		surname := individual.Name().Surname()
		if surnameStartsWith(individual, c.selectedLetter) && !stringInSlice(surnames, surname) {
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

	return html.NewRow(html.NewColumn(
		html.EntireRow, html.NewDiv("", newNavPills(pills)),
	)).String()
}
