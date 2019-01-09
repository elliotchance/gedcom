package html

import (
	"io"
)

type Anchor struct {
	name string
}

func NewAnchor(name string) *Anchor {
	return &Anchor{
		name: name,
	}
}

func (c *Anchor) WriteTo(w io.Writer) (int64, error) {
	return writeSprintf(w, `<a name="%s"/>`, c.name)
}
