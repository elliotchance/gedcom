package core

import (
	"fmt"
	"io"
)

// NavItem is a single tab in the tab bar.
type NavItem struct {
	content  Component
	href     string
	isActive bool
}

func NewNavItem(content Component, isActive bool, href string) *NavItem {
	return &NavItem{
		content:  content,
		isActive: isActive,
		href:     href,
	}
}

func (c *NavItem) WriteHTMLTo(w io.Writer) (int64, error) {
	active := activeIf(c.isActive)

	return NewTag("li", map[string]string{
		"class": "nav-item",
	}, NewTag("a", map[string]string{
		"class": fmt.Sprintf("nav-link %s", active),
		"href":  c.href,
	}, c.content)).WriteHTMLTo(w)
}
