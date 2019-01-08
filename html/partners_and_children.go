package html

import (
	"github.com/elliotchance/gedcom"
	"io"
)

// PartnersAndChildren show the partners and/or children connected to the
// individual on their individual page.
type PartnersAndChildren struct {
	individual *gedcom.IndividualNode
	document   *gedcom.Document
}

func NewPartnersAndChildren(document *gedcom.Document, individual *gedcom.IndividualNode) *PartnersAndChildren {
	return &PartnersAndChildren{
		individual: individual,
		document:   document,
	}
}

func (c *PartnersAndChildren) WriteTo(w io.Writer) (int64, error) {
	column := NewColumn(EntireRow,
		NewHeading(2, "", NewText("Spouses & Children")))

	rows := []Component{
		NewRow(column),
	}

	// Find children of known spouses.
	spouses := c.individual.Spouses()

	for _, spouse := range spouses {
		rows = append(rows, NewHorizontalRuleRow())

		columns := []*Column{
			NewColumn(QuarterRow, NewIndividualButton(c.document, spouse)),
		}

		family := c.individual.FamilyWithSpouse(spouse)
		if family != nil {
			columns, rows = partnerSection(family, c, columns, rows)
		}

		rows = append(rows,
			NewRow(columns...),
			NewRow(NewColumn(EntireRow, NewSpace())))
	}

	// Find children belonging to families with an unknown spouse.
	for _, family := range c.individual.Families() {
		// Ignore families with this individual as a child or where spouse is
		// present (since they have been handled above).
		if family.HasChild(c.individual) ||
			(family.Husband() != nil && family.Wife() != nil) {
			continue
		}

		rows = append(rows, NewHorizontalRuleRow())

		columns := []*Column{
			NewColumn(QuarterRow, NewIndividualButton(c.document, nil)),
		}

		columns, rows = partnerSection(family, c, columns, rows)

		rows = append(rows,
			NewRow(columns...),
			NewRow(NewColumn(EntireRow, NewSpace())))
	}

	if len(rows) == 1 {
		rows = append(rows,
			NewHorizontalRuleRow(),
			NewText("There are no known spouses or children."),
			NewRow(NewColumn(EntireRow, NewSpace())),
		)
	}

	return NewComponents(rows...).WriteTo(w)
}

func partnerSection(family *gedcom.FamilyNode, c *PartnersAndChildren, columns []*Column, rows []Component) ([]*Column, []Component) {
	children := family.Children()
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

		button := NewComponents(
			svg,
			NewIndividualButton(c.document, child),
		)
		columns = append(columns, NewColumn(3, button))

		if len(columns) == 4 {
			rows = append(rows, NewRow(columns...))
			columns = []*Column{
				NewColumn(QuarterRow, NewEmpty()),
			}
		}
	}

	return columns, rows
}
