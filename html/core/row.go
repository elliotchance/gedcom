package core

import "io"

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

func (c *Row) WriteHTMLTo(w io.Writer) (int64, error) {
	n := appendString(w, `<div class="row">`)

	for _, column := range c.columns {
		n += appendComponent(w, column)
	}

	n += appendString(w, "</div>")

	return n, nil
}
