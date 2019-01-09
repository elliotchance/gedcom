package html

import "io"

type LineBreak struct{}

func NewLineBreak() *LineBreak {
	return &LineBreak{}
}

func (c *LineBreak) WriteTo(w io.Writer) (int64, error) {
	return writeString(w, `<br/>`)
}
