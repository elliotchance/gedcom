package main

import "fmt"

// components is a wrapper for zero more components that rendered at the same
// time.
type components struct {
	items []fmt.Stringer
}

func newComponents(items ...fmt.Stringer) *components {
	return &components{
		items: items,
	}
}

func (c *components) String() string {
	s := ""
	for _, item := range c.items {
		s += item.String()
	}

	return s
}
