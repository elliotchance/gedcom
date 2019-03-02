package core

import (
	"io"
)

// KeyedTableRow is a table row consisting of two columns where the left column
// is a header and a key for the data in the right column. It also allows the
// row to be hidden altogether if needed.
type KeyedTableRow struct {
	title   string
	visible bool
	value   Component
}

func NewKeyedTableRow(title string, value Component, visible bool) *KeyedTableRow {
	return &KeyedTableRow{
		title:   title,
		value:   value,
		visible: visible,
	}
}

func (c *KeyedTableRow) WriteHTMLTo(w io.Writer) (int64, error) {
	if !c.visible {
		return writeNothing()
	}

	return NewComponents(
		NewTableRow(
			NewTableCell(NewText(c.title)).Header(),
			NewTableCell(c.value),
		),
	).WriteHTMLTo(w)
}
