package main

import "fmt"

// tableHead is the <thead> section of a table that contains the table heading
// cells.
type tableHead struct {
	columns []string
}

func newTableHead(columns ...string) *tableHead {
	return &tableHead{
		columns: columns,
	}
}

func (c *tableHead) String() string {
	s := `<thead><tr>`

	for _, column := range c.columns {
		s += fmt.Sprintf(`<th scope="col">%s</th>`, column)
	}

	return s + `</tr></thead>`
}
