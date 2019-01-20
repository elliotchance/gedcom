package html

import (
	"github.com/elliotchance/gedcom"
	"io"
)

// Components is a wrapper for zero more components that rendered at the same
// time.
type Components struct {
	items []Component
}

func NewComponents(items ...Component) *Components {
	nonNilItems := []Component{}

	for _, item := range items {
		if !gedcom.IsNil(item) {
			nonNilItems = append(nonNilItems, item)
		}
	}

	return &Components{
		items: nonNilItems,
	}
}

func (c *Components) WriteTo(w io.Writer) (int64, error) {
	n := int64(0)
	for _, item := range c.items {
		n += appendComponent(w, item)
	}

	return n, nil
}
