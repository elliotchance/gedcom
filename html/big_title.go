package html

type BigTitle struct {
	text string
	size int
}

func NewBigTitle(size int, text string) *BigTitle {
	return &BigTitle{
		text: text,
		size: size,
	}
}

func (c *BigTitle) String() string {
	return NewRow(
		NewColumn(EntireRow,
			NewHeading(c.size, "text-center", c.text),
		),
	).String()
}
