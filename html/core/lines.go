package core

import (
	"io"
)

type Lines struct {
	lines []Component
}

func NewLines(lines ...Component) *Lines {
	return &Lines{
		lines: lines,
	}
}

func (c *Lines) WriteHTMLTo(w io.Writer) (int64, error) {
	components := []Component{}

	for i, line := range c.lines {
		components = append(components, line)

		if i < len(c.lines)-1 {
			components = append(components, NewLineBreak())
		}
	}

	return NewComponents(components...).WriteHTMLTo(w)
}
