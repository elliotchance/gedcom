package html

import "fmt"

type StyledTableCell struct {
	content fmt.Stringer
	class   string
	style   string
}

func NewStyledTableCell(style, class string, content fmt.Stringer) *StyledTableCell {
	return &StyledTableCell{
		content: content,
		style:   style,
		class:   class,
	}
}

func (c *StyledTableCell) String() string {
	return fmt.Sprintf(`<td scope="col" class="%s" style="%s">%s</td>`,
		c.class, c.style, c.content)
}
