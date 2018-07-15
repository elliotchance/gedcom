package main

import (
	"fmt"
)

const (
	quarterRow = 3
	halfRow    = 6
	entireRow  = 12
)

// row is a page row for Bootstrap.
type row struct {
	columns []*column
}

func newRow(columns ...*column) *row {
	return &row{
		columns: columns,
	}
}

func (c *row) String() string {
	columns := ""
	for _, column := range c.columns {
		columns += column.String()
	}

	return fmt.Sprintf(`
		<div class="row">
			%s
		</div>`, columns)
}
