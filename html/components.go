package html

import (
	"io"
)

// Components is a wrapper for zero more components that rendered at the same
// time.
type Components struct {
	items []Component
}

func NewComponents(items ...Component) *Components {
	return &Components{
		items: items,
	}
}

func (c *Components) WriteTo(w io.Writer) (int64, error) {
	n := int64(0)
	for _, item := range c.items {
		n += appendComponent(w, item)
	}

	return n, nil
}
