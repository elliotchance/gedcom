package html

import "io"

// TableHead is the <thead> section of a table that contains the table heading
// cells.
type TableHead struct {
	columns []string
}

func NewTableHead(columns ...string) *TableHead {
	return &TableHead{
		columns: columns,
	}
}

func (c *TableHead) WriteTo(w io.Writer) (int64, error) {
	n := appendString(w, `<thead><tr>`)

	for _, column := range c.columns {
		n += appendSprintf(w, `<th scope="col">%s</th>`, column)
	}

	n += appendString(w, `</tr></thead>`)

	return n, nil
}
