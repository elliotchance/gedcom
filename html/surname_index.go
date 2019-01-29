package html

import (
	"github.com/elliotchance/gedcom"
	"io"
)

type SurnameIndex struct {
	document       *gedcom.Document
	selectedLetter rune
	visibility     LivingVisibility
}

func NewSurnameIndex(document *gedcom.Document, selectedLetter rune, visibility LivingVisibility) *SurnameIndex {
	return &SurnameIndex{
		document:       document,
		selectedLetter: selectedLetter,
		visibility:     visibility,
	}
}

func (c *SurnameIndex) WriteTo(w io.Writer) (int64, error) {
	surnames := gedcom.NewStringSet()

	for _, individual := range c.document.Individuals() {
		if individual.IsLiving() {
			switch c.visibility {
			case LivingVisibilityHide, LivingVisibilityPlaceholder:
				continue

			case LivingVisibilityShow:
				// Proceed.
			}
		}

		surname := individual.Name().Surname()
		if surnameStartsWith(individual, c.selectedLetter) {
			surnames.Add(surname)
		}
	}

	// Render
	pills := []Component{}
	for _, surname := range surnames.Strings() {
		pills = append(pills, NewNavLink(surname, "#"+surname, false))
	}

	return NewNavPillsRow(pills).WriteTo(w)
}
