package html

import "fmt"

// Components is a wrapper for zero more components that rendered at the same
// time.
type Components struct {
	items []fmt.Stringer
}

func NewComponents(items ...fmt.Stringer) *Components {
	return &Components{
		items: items,
	}
}

func (c *Components) String() string {
	s := ""
	for _, item := range c.items {
		s += item.String()
	}

	return s
}
