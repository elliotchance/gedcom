package core

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

func (c *TableHead) WriteHTMLTo(w io.Writer) (int64, error) {
	mutN := appendString(w, `<thead><tr>`)

	for _, column := range c.columns {
		mutN += appendSprintf(w, `<th scope="col">%s</th>`, column)
	}

	mutN += appendString(w, `</tr></thead>`)

	return mutN, nil
}
