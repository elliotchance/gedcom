package main

import "github.com/elliotchance/gedcom/html"

// keyedTableRow is a table row consisting of two columns where the left column
// is a header and a key for the data in the right column. It also allows the
// row to be hidden altogether if needed.
type keyedTableRow struct {
	title, value string
	visible      bool
}

func newKeyedTableRow(title, value string, visible bool) *keyedTableRow {
	return &keyedTableRow{
		title:   title,
		value:   value,
		visible: visible,
	}
}

func (c *keyedTableRow) String() string {
	if !c.visible {
		return ""
	}

	return html.Sprintf(`
       <tr>
           <th scope="row">%s</th>
           <td>%s</td>
       </tr>`, c.title, c.value)
}
