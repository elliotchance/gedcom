package core

import (
	"io"
)

type TableRow struct {
	cells []Component
}

func NewTableRow(cells ...Component) *TableRow {
	return &TableRow{
		cells: cells,
	}
}

func (c *TableRow) WriteHTMLTo(w io.Writer) (int64, error) {
	n := appendString(w, `<tr>`)

	for _, cell := range c.cells {
		n += appendComponent(w, cell)
	}

	n += appendString(w, `</tr>`)

	return n, nil
}
