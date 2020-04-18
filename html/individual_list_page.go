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
	placesMap         map[string]*place
}

func NewIndividualListPage(document *gedcom.Document, selectedLetter rune, googleAnalyticsID string, options *PublishShowOptions, indexLetters []rune, placesMap map[string]*place) *IndividualListPage {
	return &IndividualListPage{
		document:          document,
		selectedLetter:    selectedLetter,
		googleAnalyticsID: googleAnalyticsID,
		options:           options,
		indexLetters:      indexLetters,
		placesMap:         placesMap,
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

		table = append(table, NewIndividualInList(c.document, i,
			c.options.LivingVisibility, c.placesMap))
	}

	livingRow := core.NewRow(
		core.NewColumn(core.EntireRow, core.NewText(fmt.Sprintf(
			"%d individuals are hidden because they are living.",
			livingCount,
		))),
	)

	if livingCount == 0 ||
		c.options.LivingVisibility == LivingVisibilityHide ||
		c.options.LivingVisibility == LivingVisibilityShow {
		livingRow = nil
	}

	return core.NewPage("Individuals", core.NewComponents(
		NewPublishHeader(c.document, "", selectedIndividualsTab,
			c.options, c.indexLetters, c.placesMap),
		livingRow,
		core.NewSpace(),
		NewIndividualIndexHeader(c.document, c.selectedLetter,
			c.options.LivingVisibility, c.indexLetters),
		core.NewSpace(),
		NewSurnameIndex(c.document, c.selectedLetter, c.options.LivingVisibility),
		core.NewSpace(),
		core.NewRow(
			core.NewColumn(core.EntireRow, core.NewTable("", table...)),
		),
	), c.googleAnalyticsID).WriteHTMLTo(w)
}
