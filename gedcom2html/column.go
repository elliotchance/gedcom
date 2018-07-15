package main

import (
	"fmt"
)

// column is used inside of a row. The row consists of 12 virtual columns and
// each column can specify how many of the columns it represents.
type column struct {
	width int
	body  fmt.Stringer
}

func newColumn(width int, body fmt.Stringer) *column {
	return &column{
		width: width,
		body:  body,
	}
}

func (c *column) String() string {
	return newDiv(fmt.Sprintf("col-%d", c.width), c.body).String()
}
