package html

import "fmt"

type Lines struct {
	lines []fmt.Stringer
}

func NewLines(lines ...fmt.Stringer) *Lines {
	return &Lines{
		lines: lines,
	}
}

func (c *Lines) String() string {
	components := []fmt.Stringer{}

	for i, line := range c.lines {
		components = append(components, line)

		if i%2 == 1 {
			components = append(components, NewLineBreak())
		}
	}

	return NewComponents(components...).String()
}
