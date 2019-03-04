package core

import "io"

// Space is an empty row used as a white space separator between other page
// rows.
type Space struct{}

func NewSpace() *Space {
	return &Space{}
}

func (c *Space) WriteHTMLTo(w io.Writer) (int64, error) {
	return NewRow(NewColumn(EntireRow, NewHTML("&nbsp;"))).WriteHTMLTo(w)
}
