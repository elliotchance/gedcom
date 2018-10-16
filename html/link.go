package html

import "fmt"

type Link struct {
	text  string
	dest  string
	style string
}

func NewLink(text, dest string) *Link {
	return &Link{
		text: text,
		dest: dest,
	}
}

func (c *Link) Style(style string) *Link {
	c.style = style

	return c
}

func (c *Link) String() string {
	attributes := ""
	if c.style != "" {
		attributes += fmt.Sprintf(` style="%s"`, c.style)
	}

	return fmt.Sprintf(`<a href="%s"%s>%s</a>`, c.dest, attributes, c.text)
}
