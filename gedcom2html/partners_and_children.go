package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
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
		newRow(newColumn(entireRow, newHeading(2, "", "Spouses & Children"))),
	}

	// Find children of known spouses.
	spouses := c.individual.Spouses(c.document)

	for _, spouse := range spouses {
		rows = append(rows, newHorizontalRuleRow())

		columns := []*column{
			newColumn(quarterRow, newIndividualButton(spouse)),
		}

		family := c.individual.FamilyWithSpouse(c.document, spouse)
		if family != nil {
			columns, rows = partnerSection(family, c, columns, rows)
		}

		rows = append(rows,
			newRow(columns...),
			newRow(newColumn(entireRow, newSpace())))
	}

	// Find children belonging to families with an unknown spouse.
	for _, family := range c.individual.Families(c.document) {
		// Ignore families with this individual as a child or where spouse is
		// present (since they have been handled above).
		if family.HasChild(c.document, c.individual) ||
			(family.Husband(c.document) != nil && family.Wife(c.document) != nil) {
			continue
		}

		rows = append(rows, newHorizontalRuleRow())

		columns := []*column{
			newColumn(quarterRow, newIndividualButton(nil)),
		}

		columns, rows = partnerSection(family, c, columns, rows)

		rows = append(rows,
			newRow(columns...),
			newRow(newColumn(entireRow, newSpace())))
	}

	if len(rows) == 1 {
		rows = append(rows,
			newHorizontalRuleRow(),
			newText("There are no known spouses or children."),
			newRow(newColumn(entireRow, newSpace())),
		)
	}

	return newComponents(rows...).String()
}

func partnerSection(family *gedcom.FamilyNode, c *partnersAndChildren, columns []*column, rows []fmt.Stringer) ([]*column, []fmt.Stringer) {
	children := family.Children(c.document)
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

		button := newComponents(
			svg,
			newIndividualButton(child),
		)
		columns = append(columns, newColumn(3, button))

		if len(columns) == 4 {
			rows = append(rows, newRow(columns...))
			columns = []*column{
				newColumn(quarterRow, newEmpty()),
			}
		}
	}

	return columns, rows
}
