package html

import (
	"io"
)

type StyledTableCell struct {
	content Component
	class   string
	style   string
}

func NewStyledTableCell(style, class string, content Component) *StyledTableCell {
	return &StyledTableCell{
		content: content,
		style:   style,
		class:   class,
	}
}

func (c *StyledTableCell) WriteTo(w io.Writer) (int64, error) {
	attributes := map[string]string{
		"scope": "col",
		"class": c.class,
		"style": c.style,
	}

	return NewTag("td", attributes, c.content).WriteTo(w)
}
