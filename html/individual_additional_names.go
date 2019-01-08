package html

import (
	"github.com/elliotchance/gedcom"
	"io"
)

// IndividualAdditionalNames is shown on the individual page. It shows all of
// the extra names (except the primary name) and their type.
type IndividualAdditionalNames struct {
	individual *gedcom.IndividualNode
}

func NewIndividualAdditionalNames(individual *gedcom.IndividualNode) *IndividualAdditionalNames {
	return &IndividualAdditionalNames{
		individual: individual,
	}
}

func (c *IndividualAdditionalNames) WriteTo(w io.Writer) (int64, error) {
	rows := []Component{}
	names := c.individual.Names()

	for _, name := range names {
		row := NewKeyedTableRow(name.Type().String(), NewText(name.String()),
			name.Type() != "")
		rows = append(rows, row)
	}

	table := NewTable("", rows...)

	return NewCard("Additional Names", len(names)-1, table).WriteTo(w)
}
