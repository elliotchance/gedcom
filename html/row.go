package html

import (
	"fmt"
)

const (
	QuarterRow = 3
	HalfRow    = 6
	EntireRow  = 12
)

// Row is a page row for Bootstrap.
type Row struct {
	columns []*Column
}

func NewRow(columns ...*Column) *Row {
	return &Row{
		columns: columns,
	}
}

func (c *Row) String() string {
	columns := ""
	for _, column := range c.columns {
		columns += column.String()
	}

	return fmt.Sprintf(`
		<div class="row">
			%s
		</div>`, columns)
}
