package html

import "fmt"

type TableRow struct {
	cells []fmt.Stringer
}

func NewTableRow(cells ...fmt.Stringer) *TableRow {
	return &TableRow{
		cells: cells,
	}
}

func (c *TableRow) String() string {
	s := `<tr>`

	for _, cell := range c.cells {
		s += cell.String()
	}

	return s + `</tr>`
}
