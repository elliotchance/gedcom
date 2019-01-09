package html

import (
	"io"
)

type NavPillsRow struct {
	links []Component
}

func NewNavPillsRow(links []Component) *NavPillsRow {
	return &NavPillsRow{
		links: links,
	}
}

func (c *NavPillsRow) WriteTo(w io.Writer) (int64, error) {
	pills := NewNavPills(c.links)
	div := NewDiv("", pills)
	column := NewColumn(EntireRow, div)
	row := NewRow(column)

	return row.WriteTo(w)
}
