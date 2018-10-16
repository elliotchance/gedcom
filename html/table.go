package html

import "fmt"

// Table is a HTML table.
type Table struct {
	content    []fmt.Stringer
	tableClass string
}

func NewTable(tableClass string, content ...fmt.Stringer) *Table {
	return &Table{
		content:    content,
		tableClass: tableClass,
	}
}

func (c *Table) String() string {
	components := NewComponents(c.content...).String()

	return Sprintf(`<table class="table %s">%s</table>`,
		c.tableClass, components)
}
