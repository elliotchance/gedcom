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
	options           *PublishShowOptions
	indexLetters      []rune
}

func NewIndividualListPage(document *gedcom.Document, selectedLetter rune, googleAnalyticsID string, options *PublishShowOptions, indexLetters []rune) *IndividualListPage {
	return &IndividualListPage{
		document:          document,
		selectedLetter:    selectedLetter,
		googleAnalyticsID: googleAnalyticsID,
		options:           options,
		indexLetters:      indexLetters,
	}
}

func (c *IndividualListPage) WriteHTMLTo(w io.Writer) (int64, error) {
	table := []core.Component{
		core.NewTableHead("Name", "Birth", "Death"),
	}

	mutIndividuals := gedcom.IndividualNodes{}

	for _, individual := range c.document.Individuals() {
		if surnameStartsWith(individual, c.selectedLetter) {
			mutIndividuals = append(mutIndividuals, individual)
		}
	}

	// Sort mutIndividuals by name.
	sort.Slice(mutIndividuals, func(i, j int) bool {
		left := mutIndividuals[i].Name().Format(gedcom.NameFormatIndex)
		right := mutIndividuals[j].Name().Format(gedcom.NameFormatIndex)

		return left < right
	})

	livingCount := 0
	lastSurname := ""
	for _, i := range mutIndividuals {
		if i.IsLiving() {
			switch c.options.LivingVisibility {
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

		table = append(table, NewIndividualInList(c.document, i, c.options.LivingVisibility))
	}

	livingRow := core.NewRow(
		core.NewColumn(core.EntireRow, core.NewText(fmt.Sprintf(
			"%d mutIndividuals are hidden because they are living.",
			livingCount,
		))),
	)

	if livingCount == 0 ||
		c.options.LivingVisibility == LivingVisibilityHide ||
		c.options.LivingVisibility == LivingVisibilityShow {
		livingRow = nil
	}

	return core.NewPage("Individuals", core.NewComponents(
		NewPublishHeader(c.document, "", selectedIndividualsTab, c.options, c.indexLetters),
		livingRow,
		core.NewSpace(),
		NewIndividualIndexHeader(c.document, c.selectedLetter, c.options.LivingVisibility, c.indexLetters),
		core.NewSpace(),
		NewSurnameIndex(c.document, c.selectedLetter, c.options.LivingVisibility),
		core.NewSpace(),
		core.NewRow(
			core.NewColumn(core.EntireRow, core.NewTable("", table...)),
		),
	), c.googleAnalyticsID).WriteHTMLTo(w)
}
