package html

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"io"
	"sort"
)

// IndividualListPage is the page that lists of all the individuals.
type IndividualListPage struct {
	document          *gedcom.Document
	selectedLetter    rune
	googleAnalyticsID string
	options           PublishShowOptions
	visibility        LivingVisibility
}

func NewIndividualListPage(document *gedcom.Document, selectedLetter rune, googleAnalyticsID string, options PublishShowOptions, visibility LivingVisibility) *IndividualListPage {
	return &IndividualListPage{
		document:          document,
		selectedLetter:    selectedLetter,
		googleAnalyticsID: googleAnalyticsID,
		options:           options,
		visibility:        visibility,
	}
}

func (c *IndividualListPage) WriteTo(w io.Writer) (int64, error) {
	table := []Component{
		NewTableHead("Name", "Birth", "Death"),
	}

	individuals := gedcom.IndividualNodes{}

	for _, individual := range c.document.Individuals() {
		if surnameStartsWith(individual, c.selectedLetter) {
			individuals = append(individuals, individual)
		}
	}

	// Sort individuals by name.
	sort.Slice(individuals, func(i, j int) bool {
		left := individuals[i].Name().Format(gedcom.NameFormatIndex)
		right := individuals[j].Name().Format(gedcom.NameFormatIndex)

		return left < right
	})

	livingCount := 0
	lastSurname := ""
	for _, i := range individuals {
		if i.IsLiving() {
			switch c.visibility {
			case LivingVisibilityShow:
				// Proceed.

			case LivingVisibilityHide, LivingVisibilityPlaceholder:
				livingCount += 1
				continue
			}
		}

		if newSurname := i.Name().Surname(); newSurname != lastSurname {
			heading := NewComponents(
				NewAnchor(newSurname),
				NewHeading(3, "", NewText(newSurname)),
			)

			table = append(table, NewTableRow(
				NewTableCell(heading),
				NewTableCell(NewText("")),
				NewTableCell(NewText("")),
			))

			lastSurname = newSurname
		}

		table = append(table, NewIndividualInList(c.document, i, c.visibility))
	}

	livingRow := NewRow(
		NewColumn(EntireRow, NewText(fmt.Sprintf(
			"%d individuals are hidden because they are living.",
			livingCount,
		))),
	)

	if livingCount == 0 ||
		c.visibility == LivingVisibilityHide ||
		c.visibility == LivingVisibilityShow {
		livingRow = nil
	}

	return NewPage("Individuals", NewComponents(
		NewPublishHeader(c.document, "", selectedIndividualsTab, c.options),
		livingRow,
		NewSpace(),
		NewIndividualIndexHeader(c.document, c.selectedLetter),
		NewSpace(),
		NewSurnameIndex(c.document, c.selectedLetter, c.visibility),
		NewSpace(),
		NewRow(
			NewColumn(EntireRow, NewTable("", table...)),
		),
	), c.googleAnalyticsID).WriteTo(w)
}
