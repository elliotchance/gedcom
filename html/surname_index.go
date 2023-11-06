package html

import (
	"io"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/gedcom/v39/html/core"
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

func (c *SurnameIndex) WriteHTMLTo(w io.Writer) (int64, error) {
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
	pills := []core.Component{}
	for _, surname := range surnames.Strings() {
		pills = append(pills, core.NewNavLink(surname, "#"+surname, false))
	}

	return core.NewNavPillsRow(pills).WriteHTMLTo(w)
}
