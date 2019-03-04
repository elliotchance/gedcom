package html

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
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

func (c *IndividualListPage) WriteHTMLTo(w io.Writer) (int64, error) {
	table := []core.Component{
		core.NewTableHead("Name", "Birth", "Death"),
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
			heading := core.NewComponents(
				core.NewAnchor(newSurname),
				core.NewHeading(3, "", core.NewText(newSurname)),
			)

			table = append(table, core.NewTableRow(
				core.NewTableCell(heading),
				core.NewTableCell(core.NewText("")),
				core.NewTableCell(core.NewText("")),
			))

			lastSurname = newSurname
		}

		table = append(table, NewIndividualInList(c.document, i, c.visibility))
	}

	livingRow := core.NewRow(
		core.NewColumn(core.EntireRow, core.NewText(fmt.Sprintf(
			"%d individuals are hidden because they are living.",
			livingCount,
		))),
	)

	if livingCount == 0 ||
		c.visibility == LivingVisibilityHide ||
		c.visibility == LivingVisibilityShow {
		livingRow = nil
	}

	return core.NewPage("Individuals", core.NewComponents(
		NewPublishHeader(c.document, "", selectedIndividualsTab, c.options),
		livingRow,
		core.NewSpace(),
		NewIndividualIndexHeader(c.document, c.selectedLetter),
		core.NewSpace(),
		NewSurnameIndex(c.document, c.selectedLetter, c.visibility),
		core.NewSpace(),
		core.NewRow(
			core.NewColumn(core.EntireRow, core.NewTable("", table...)),
		),
	), c.googleAnalyticsID).WriteHTMLTo(w)
}
