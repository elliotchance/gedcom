package html

import (
	"io"
)

// NavTabs is a group of tabs.
type NavTabs struct {
	items []*NavItem
}

func NewNavTabs(items []*NavItem) *NavTabs {
	return &NavTabs{
		items: items,
	}
}

func (c *NavTabs) WriteTo(w io.Writer) (int64, error) {
	tabs := []Component{}
	for _, tab := range c.items {
		tabs = append(tabs, tab)
	}

	return NewRow(
		NewColumn(EntireRow,
			NewTag("ul", map[string]string{
				"class": "nav nav-tabs",
			}, NewComponents(tabs...)),
		),
	).WriteTo(w)
}
