package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
)

// individualAdditionalNames is shown on the individual page. It shows all of
// the extra names (except the primary name) and their type.
type individualAdditionalNames struct {
	individual *gedcom.IndividualNode
}

func newIndividualAdditionalNames(individual *gedcom.IndividualNode) *individualAdditionalNames {
	return &individualAdditionalNames{
		individual: individual,
	}
}

func (c *individualAdditionalNames) String() string {
	rows := []fmt.Stringer{}
	names := c.individual.Names()

	for _, name := range names {
		rows = append(rows, newKeyedTableRow(name.Type().String(), name.String(), name.Type() != ""))
	}

	return newCard("Additional Names", len(names)-1,
		newTable("", rows...)).String()
}
