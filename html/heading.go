package html

// Heading is larger text.
type Heading struct {
	text, class string
	number      int
}

func NewHeading(number int, class, text string) *Heading {
	return &Heading{
		text:   text,
		number: number,
		class:  class,
	}
}

func (c *Heading) String() string {
	return Sprintf(`<h%d class="%s">%s</h%d>`,
		c.number, c.class, c.text, c.number)
}
