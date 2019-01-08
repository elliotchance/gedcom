package html

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

type Tag struct {
	tag        string
	attributes map[string]string
	body       Component
}

func NewTag(tag string, attributes map[string]string, body Component) *Tag {
	return &Tag{
		tag:        tag,
		attributes: attributes,
		body:       body,
	}
}

func (c *Tag) WriteTo(w io.Writer) (int64, error) {
	names := []string{}
	for name := range c.attributes {
		names = append(names, name)
	}

	sort.Strings(names)

	attributes := " "
	for _, name := range names {
		value := c.attributes[name]
		if value != "" {
			attributes += fmt.Sprintf(`%s="%s" `, name, value)
		}
	}

	return NewComponents(
		NewHTML(fmt.Sprintf(`<%s%s>`, c.tag, strings.TrimRight(attributes, " "))),
		c.body,
		NewHTML(fmt.Sprintf(`</%s>`, c.tag)),
	).WriteTo(w)
}
