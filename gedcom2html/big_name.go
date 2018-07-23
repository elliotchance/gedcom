package main

// bigName shows the large name on the top of a page.
type bigName struct {
	text string
}

func newBigName(text string) *bigName {
	return &bigName{
		text: text,
	}
}

func (c *bigName) String() string {
	return newRow(
		newColumn(entireRow,
			newHeading(1, "text-center", c.text),
		),
	).String()
}
