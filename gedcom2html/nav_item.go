package main

import (
	"fmt"
)

// navItem is a single tab in the tab bar.
type navItem struct {
	content, href string
	isActive      bool
}

func newNavItem(content string, isActive bool, href string) *navItem {
	return &navItem{
		content:  content,
		isActive: isActive,
		href:     href,
	}
}

func (c *navItem) String() string {
	active := ""
	if c.isActive {
		active = "active"
	}

	return fmt.Sprintf(`
		<li class="nav-item">
			<a class="nav-link %s" href="%s">%s</a>
		</li>`, active, c.href, c.content)
}
