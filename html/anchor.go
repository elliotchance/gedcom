package html

import "fmt"

type Anchor struct {
	name string
}

func NewAnchor(name string) *Anchor {
	return &Anchor{
		name: name,
	}
}

func (c *Anchor) String() string {
	return fmt.Sprintf(`<a name="%s"/>`, c.name)
}
