package core

import (
	"fmt"
	"io"
)

type NavLink struct {
	text       string
	link       string
	isSelected bool
}

func NewNavLink(text, link string, isSelected bool) *NavLink {
	return &NavLink{
		text:       text,
		link:       link,
		isSelected: isSelected,
	}
}

func (c *NavLink) WriteHTMLTo(w io.Writer) (int64, error) {
	active := ""
	if c.isSelected {
		active = "active"
	}

	return NewTag("li", map[string]string{
		"class": "nav-item",
	}, NewTag("a", map[string]string{
		"class": fmt.Sprintf("nav-link %s", active),
		"href":  c.link,
	}, NewText(c.text))).WriteHTMLTo(w)
}
