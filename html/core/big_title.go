package core

import "io"

type BigTitle struct {
	text Component
	size int
}

func NewBigTitle(size int, text Component) *BigTitle {
	return &BigTitle{
		text: text,
		size: size,
	}
}

func (c *BigTitle) WriteHTMLTo(w io.Writer) (int64, error) {
	return NewRow(
		NewColumn(EntireRow,
			NewHeading(c.size, "text-center", c.text),
		),
	).WriteHTMLTo(w)
}
