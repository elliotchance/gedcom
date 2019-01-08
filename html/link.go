package html

import (
	"io"
)

type Link struct {
	body  Component
	dest  string
	style string
}

func NewLink(body Component, dest string) *Link {
	return &Link{
		body: body,
		dest: dest,
	}
}

func (c *Link) Style(style string) *Link {
	c.style = style

	return c
}

func (c *Link) WriteTo(w io.Writer) (int64, error) {
	attributes := map[string]string{
		"style": c.style,
		"href":  c.dest,
	}

	return NewTag("a", attributes, c.body).WriteTo(w)
}
