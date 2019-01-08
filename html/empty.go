package html

import "io"

// Empty is used a placeholder for a component where nothing should be visible.
type Empty struct{}

func NewEmpty() *Empty {
	return &Empty{}
}

func (c *Empty) WriteTo(w io.Writer) (int64, error) {
	return writeString(w, "&nbsp;")
}
