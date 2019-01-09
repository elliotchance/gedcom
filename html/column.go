package html

import (
	"fmt"
	"io"
)

// Column is used inside of a row. The row consists of 12 virtual columns and
// each column can specify how many of the columns it represents.
type Column struct {
	width int
	body  Component
}

func NewColumn(width int, body Component) *Column {
	return &Column{
		width: width,
		body:  body,
	}
}

func (c *Column) WriteTo(w io.Writer) (int64, error) {
	return NewDiv(fmt.Sprintf("col-%d", c.width), c.body).WriteTo(w)
}
