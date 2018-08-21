package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
)

// partnersAndChildren show the partners and/or children connected to the
// individual on their individual page.
type partnersAndChildren struct {
	individual *gedcom.IndividualNode
	document   *gedcom.Document
}

func newPartnersAndChildren(document *gedcom.Document, individual *gedcom.IndividualNode) *partnersAndChildren {
	return &partnersAndChildren{
		individual: individual,
		document:   document,
	}
}

func (c *partnersAndChildren) String() string {
	rows := []fmt.Stringer{
		html.NewRow(html.NewColumn(html.EntireRow,
			html.NewHeading(2, "", "Spouses & Children"))),
	}

	// Find children of known spouses.
	spouses := c.individual.Spouses()

	for _, spouse := range spouses {
		rows = append(rows, html.NewHorizontalRuleRow())

		columns := []*html.Column{
			html.NewColumn(html.QuarterRow, newIndividualButton(c.document, spouse)),
		}

		family := c.individual.FamilyWithSpouse(spouse)
		if family != nil {
			columns, rows = partnerSection(family, c, columns, rows)
		}

		rows = append(rows,
			html.NewRow(columns...),
			html.NewRow(html.NewColumn(html.EntireRow, html.NewSpace())))
	}

	// Find children belonging to families with an unknown spouse.
	for _, family := range c.individual.Families() {
		// Ignore families with this individual as a child or where spouse is
		// present (since they have been handled above).
		if family.HasChild(c.individual) ||
			(family.Husband() != nil && family.Wife() != nil) {
			continue
		}

		rows = append(rows, html.NewHorizontalRuleRow())

		columns := []*html.Column{
			html.NewColumn(html.QuarterRow, newIndividualButton(c.document, nil)),
		}

		columns, rows = partnerSection(family, c, columns, rows)

		rows = append(rows,
			html.NewRow(columns...),
			html.NewRow(html.NewColumn(html.EntireRow, html.NewSpace())))
	}

	if len(rows) == 1 {
		rows = append(rows,
			html.NewHorizontalRuleRow(),
			html.NewText("There are no known spouses or children."),
			html.NewRow(html.NewColumn(html.EntireRow, html.NewSpace())),
		)
	}

	return html.NewComponents(rows...).String()
}

func partnerSection(family *gedcom.FamilyNode, c *partnersAndChildren, columns []*html.Column, rows []fmt.Stringer) ([]*html.Column, []fmt.Stringer) {
	children := family.Children()
	numberOfChildren := len(children)
	for i, child := range children {
		svg := newPlusSVG(false, true, true, true)

		if i > 2 {
			// These will be all of the children in the second row.
			svg = newPlusSVG(true, false, false, true)
		}

		if i == 2 || (i == numberOfChildren-1 && i < 3) {
			// If this is the last child on the first row.
			svg = newPlusSVG(false, true, false, true)
		}

		button := html.NewComponents(
			svg,
			newIndividualButton(c.document, child),
		)
		columns = append(columns, html.NewColumn(3, button))

		if len(columns) == 4 {
			rows = append(rows, html.NewRow(columns...))
			columns = []*html.Column{
				html.NewColumn(html.QuarterRow, newEmpty()),
			}
		}
	}

	return columns, rows
}
