package html

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
	"io"
)

// PartnersAndChildren show the partners and/or children connected to the
// individual on their individual page.
type PartnersAndChildren struct {
	individual *gedcom.IndividualNode
	document   *gedcom.Document
	visibility LivingVisibility
	placesMap  map[string]*place
}

func NewPartnersAndChildren(document *gedcom.Document, individual *gedcom.IndividualNode, visibility LivingVisibility, placesMap map[string]*place) *PartnersAndChildren {
	return &PartnersAndChildren{
		individual: individual,
		document:   document,
		visibility: visibility,
		placesMap:  placesMap,
	}
}

func (c *PartnersAndChildren) WriteHTMLTo(w io.Writer) (int64, error) {
	heading := core.NewHeading(2, "", core.NewText("Spouses & Children"))
	column := core.NewColumn(core.EntireRow, heading)

	rows := []core.Component{
		core.NewRow(column),
	}

	// Find children of known spouses.
	spouses := c.individual.Spouses()

	for _, spouse := range spouses {
		if spouse.IsLiving() {
			switch c.visibility {
			case LivingVisibilityHide:
				continue

			case LivingVisibilityShow, LivingVisibilityPlaceholder:
				// Proceed.
			}
		}

		rows = append(rows, core.NewHorizontalRuleRow())

		columns := []*core.Column{
			core.NewColumn(core.QuarterRow, NewIndividualButton(c.document,
				spouse, c.visibility, c.placesMap)),
		}

		family := c.individual.FamilyWithSpouse(spouse)
		if family != nil {
			columns, rows = partnerSection(family, c, columns, rows)
		}

		rows = append(rows,
			core.NewRow(columns...),
			core.NewRow(core.NewColumn(core.EntireRow, core.NewSpace())))
	}

	// Find children belonging to families with an unknown spouse.
	for _, family := range c.individual.Families() {
		// Ignore families with this individual as a child or where spouse is
		// present (since they have been handled above).
		if family.HasChild(c.individual) ||
			(family.Husband() != nil && family.Wife() != nil) {
			continue
		}

		rows = append(rows, core.NewHorizontalRuleRow())

		columns := []*core.Column{
			core.NewColumn(core.QuarterRow,
				NewIndividualButton(c.document, nil, c.visibility,
					c.placesMap)),
		}

		columns, rows = partnerSection(family, c, columns, rows)

		rows = append(rows,
			core.NewRow(columns...),
			core.NewRow(core.NewColumn(core.EntireRow, core.NewSpace())))
	}

	if len(rows) == 1 {
		rows = append(rows,
			core.NewHorizontalRuleRow(),
			core.NewText("There are no known spouses or children."),
			core.NewRow(core.NewColumn(core.EntireRow, core.NewSpace())),
		)
	}

	return core.NewComponents(rows...).WriteHTMLTo(w)
}

func partnerSection(family *gedcom.FamilyNode, c *PartnersAndChildren, columns []*core.Column, rows []core.Component) ([]*core.Column, []core.Component) {
	allChildren := family.Children()
	children := []*gedcom.IndividualNode{}

	for _, child := range allChildren {
		if child.Individual().IsLiving() {
			switch c.visibility {
			case LivingVisibilityHide:
				continue

			case LivingVisibilityShow, LivingVisibilityPlaceholder:
				// Proceed.
			}
		}

		children = append(children, child.Individual())
	}

	numberOfChildren := len(children)
	for i, child := range children {
		svg := NewPlusSVG(false, true, true, true)

		if i > 2 {
			// These will be all of the children in the second row.
			svg = NewPlusSVG(true, false, false, true)
		}

		if i == 2 || (i == numberOfChildren-1 && i < 3) {
			// If this is the last child on the first row.
			svg = NewPlusSVG(false, true, false, true)
		}

		button := core.NewComponents(
			svg,
			NewIndividualButton(c.document, child, c.visibility, c.placesMap),
		)
		columns = append(columns, core.NewColumn(3, button))

		if len(columns) == 4 {
			rows = append(rows, core.NewRow(columns...))
			columns = []*core.Column{
				core.NewColumn(core.QuarterRow, core.NewEmpty()),
			}
		}
	}

	return columns, rows
}
