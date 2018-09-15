package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
	"sort"
	"strings"
)

// individualListPage is the page that lists of all the individuals.
type individualListPage struct {
	document          *gedcom.Document
	selectedLetter    rune
	googleAnalyticsID string
}

func newIndividualListPage(document *gedcom.Document, selectedLetter rune, googleAnalyticsID string) *individualListPage {
	return &individualListPage{
		document:          document,
		selectedLetter:    selectedLetter,
		googleAnalyticsID: googleAnalyticsID,
	}
}

func (c *individualListPage) String() string {
	table := []fmt.Stringer{
		html.NewTableHead("Name", "Birth", "Death"),
	}

	individuals := gedcom.IndividualNodes{}

	for _, individual := range c.document.Individuals() {
		name := individual.Name().Format(gedcom.NameFormatIndex)
		if name == "" {
			name = "#"
		}

		lowerName := strings.ToLower(name)
		if rune(lowerName[0]) == c.selectedLetter {
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
	for _, i := range individuals {
		if i.IsLiving() {
			livingCount += 1
			continue
		}

		table = append(table, newIndividualInList(c.document, i))
	}

	return html.NewPage("Individuals", html.NewComponents(
		newHeader(c.document, "", selectedIndividualsTab),
		html.NewRow(
			html.NewColumn(html.EntireRow, html.NewText(html.Sprintf(
				"%d individuals are hidden because they are living.",
				livingCount,
			))),
		),
		html.NewSpace(),
		newIndividualIndexHeader(c.document, c.selectedLetter),
		html.NewSpace(),
		html.NewRow(
			html.NewColumn(html.EntireRow, html.NewTable("", table...)),
		),
	), c.googleAnalyticsID).String()
}
