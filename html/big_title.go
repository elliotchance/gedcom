package html

type BigTitle struct {
	text string
}

func NewBigTitle(text string) *BigTitle {
	return &BigTitle{
		text: text,
	}
}

func (c *BigTitle) String() string {
	return NewRow(
		NewColumn(EntireRow,
			NewHeading(1, "text-center", c.text),
		),
	).String()
}
