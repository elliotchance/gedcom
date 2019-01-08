package html

import (
	"fmt"
	"io"
)

// Heading is larger text.
type Heading struct {
	class  string
	number int
	body   Component
}

func NewHeading(number int, class string, body Component) *Heading {
	return &Heading{
		number: number,
		class:  class,
		body:   body,
	}
}

func (c *Heading) WriteTo(w io.Writer) (int64, error) {
	attributes := map[string]string{
		"class": c.class,
	}

	return NewTag(
		fmt.Sprintf(`h%d`, c.number),
		attributes,
		c.body,
	).WriteTo(w)
}
