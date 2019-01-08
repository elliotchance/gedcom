package html

import (
	"io"
)

type TableCell struct {
	content      Component
	class, style string
	noWrap       bool
	isHeader     bool
}

func NewTableCell(content Component) *TableCell {
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

func (c *TableCell) Header() *TableCell {
	c.isHeader = true

	return c
}

func (c *TableCell) WriteTo(w io.Writer) (int64, error) {
	htmlTag := "td"
	if c.isHeader {
		htmlTag = "th"
	}

	n := appendSprintf(w, `<%s scope="col"`, htmlTag)

	if c.class != "" {
		n += appendSprintf(w, ` class="%s"`, c.class)
	}

	if c.noWrap {
		n += appendString(w, ` nowrap="nowrap"`)
	}

	if c.style != "" {
		n += appendSprintf(w, ` style="%s"`, c.style)
	}

	n += appendString(w, `>`)
	n += appendComponent(w, c.content)
	n += appendSprintf(w, `</%s>`, htmlTag)

	return n, nil
}
