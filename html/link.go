package html

import "fmt"

type Link struct {
	text string
	dest string
}

func NewLink(text, dest string) *Link {
	return &Link{
		text: text,
		dest: dest,
	}
}

func (c *Link) String() string {
	return fmt.Sprintf(`<a href="%s">%s</a>`, c.dest, c.text)
}
