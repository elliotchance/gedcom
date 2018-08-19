package html

import (
	"fmt"
)

// Column is used inside of a row. The row consists of 12 virtual columns and
// each column can specify how many of the columns it represents.
type Column struct {
	width int
	body  fmt.Stringer
}

func NewColumn(width int, body fmt.Stringer) *Column {
	return &Column{
		width: width,
		body:  body,
	}
}

func (c *Column) String() string {
	return NewDiv(fmt.Sprintf("col-%d", c.width), c.body).String()
}
