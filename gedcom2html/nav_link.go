package main

import "fmt"

type navLink struct {
	text       string
	link       string
	isSelected bool
}

func newNavLink(text, link string, isSelected bool) *navLink {
	return &navLink{
		text:       text,
		link:       link,
		isSelected: isSelected,
	}
}

func (c *navLink) String() string {
	active := ""
	if c.isSelected {
		active = "active"
	}

	return fmt.Sprintf(`
			<li class="nav-item">
    			<a class="nav-link %s" href="%s">%s</a>
  			</li>`,
		active, c.link, c.text)
}
