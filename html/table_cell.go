package html

import "fmt"

type TableCell struct {
	content fmt.Stringer
	class   string
}

func NewTableCell(class string, content fmt.Stringer) *TableCell {
	return &TableCell{
		content: content,
		class:   class,
	}
}

func (c *TableCell) String() string {
	return Sprintf(`<td scope="col" class="%s">%s</td>`, c.class, c.content)
}
