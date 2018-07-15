package main

import "fmt"

// table is a HTML table.
type table struct {
	content    []fmt.Stringer
	tableClass string
}

func newTable(tableClass string, content ...fmt.Stringer) *table {
	return &table{
		content:    content,
		tableClass: tableClass,
	}
}

func (c *table) String() string {
	return fmt.Sprintf(`<table class="table %s">%s</table>`,
		c.tableClass, newComponents(c.content...).String())
}
