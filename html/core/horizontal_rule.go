package core

import "io"

// HorizontalRule is a dividing line.
type HorizontalRule struct{}

func NewHorizontalRule() *HorizontalRule {
	return &HorizontalRule{}
}

func (c *HorizontalRule) WriteHTMLTo(w io.Writer) (int64, error) {
	return writeString(w, "<hr/>")
}
