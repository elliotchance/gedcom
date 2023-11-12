package core

import (
	"io"
)

// Table is a HTML table.
type Table struct {
	content    []Component
	tableClass string
}

func NewTable(tableClass string, content ...Component) *Table {
	return &Table{
		content:    content,
		tableClass: tableClass,
	}
}

func (c *Table) WriteHTMLTo(w io.Writer) (int64, error) {
	n := appendSprintf(w, `<table class="table table-responsive %s">`, c.tableClass)
	n += appendComponent(w, NewComponents(c.content...))
	n += appendString(w, "</table>")

	return n, nil
}
