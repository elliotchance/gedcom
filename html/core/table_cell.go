package core

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

func (c *TableCell) WriteHTMLTo(w io.Writer) (int64, error) {
	htmlTag := c.htmlTag()
	mutN := appendSprintf(w, `<%s scope="col"`, htmlTag)

	if c.class != "" {
		mutN += appendSprintf(w, ` class="%s"`, c.class)
	}

	if c.noWrap {
		mutN += appendString(w, ` nowrap="nowrap"`)
	}

	if c.style != "" {
		mutN += appendSprintf(w, ` style="%s"`, c.style)
	}

	mutN += appendString(w, `>`)
	mutN += appendComponent(w, c.content)
	mutN += appendSprintf(w, `</%s>`, htmlTag)

	return mutN, nil
}

func (c *TableCell) htmlTag() string {
	if c.isHeader {
		return "th"
	}

	return "td"
}
