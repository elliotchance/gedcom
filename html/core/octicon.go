package core

import (
	"fmt"
	"io"
)

type Octicon struct {
	name  string
	style string
}

func NewOcticon(name, style string) *Octicon {
	return &Octicon{
		name:  name,
		style: style,
	}
}

func (c *Octicon) WriteHTMLTo(w io.Writer) (int64, error) {
	return NewTag("span", map[string]string{
		"class": fmt.Sprintf("Octicon Octicon-%s", c.name),
		"style": c.style,
	}, NewText("")).WriteHTMLTo(w)
}
