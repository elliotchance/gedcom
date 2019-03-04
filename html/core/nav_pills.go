package core

import (
	"io"
)

type NavPills struct {
	links []Component
}

func NewNavPills(links []Component) *NavPills {
	return &NavPills{
		links: links,
	}
}

func (c *NavPills) WriteHTMLTo(w io.Writer) (int64, error) {
	return NewTag("ul", map[string]string{
		"class": "nav nav-pills nav-fill",
	}, NewComponents(c.links...)).WriteHTMLTo(w)
}
