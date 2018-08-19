package html

import "fmt"

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

func (c *TableHead) String() string {
	s := `<thead><tr>`

	for _, column := range c.columns {
		s += fmt.Sprintf(`<th scope="col">%s</th>`, column)
	}

	return s + `</tr></thead>`
}
