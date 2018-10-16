package html

import "fmt"

type TableCell struct {
	content      fmt.Stringer
	class, style string
	noWrap       bool
}

func NewTableCell(content fmt.Stringer) *TableCell {
	return &TableCell{
		content: content,
	}
}

func (c *TableCell) NoWrap() *TableCell {
	c.noWrap = true

	return c
}

func (c *TableCell) Class(class string) *TableCell {
	c.class = class

	return c
}

func (c *TableCell) Style(style string) *TableCell {
	c.style = style

	return c
}

func (c *TableCell) String() string {
	tag := `<td scope="col"`

	if c.class != "" {
		tag += fmt.Sprintf(` class="%s"`, c.class)
	}

	if c.noWrap {
		tag += ` nowrap="nowrap"`
	}

	if c.style != "" {
		tag += fmt.Sprintf(` style="%s"`, c.style)
	}

	return fmt.Sprintf(`%s>%s</td>`, tag, c.content)
}
