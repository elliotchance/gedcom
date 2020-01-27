package html

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
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

func (c *IndividualAdditionalNames) WriteHTMLTo(w io.Writer) (int64, error) {
	rows := []core.Component{}
	names := c.individual.Names()[1:]

	for _, name := range names {
		row := core.NewKeyedTableRow(
			name.Type().String(), core.NewText(name.String()), true)
		rows = append(rows, row)
	}

	table := core.NewTable("", rows...)

	return core.NewCard(core.NewText("Additional Names"), len(names), table).
		WriteHTMLTo(w)
}
