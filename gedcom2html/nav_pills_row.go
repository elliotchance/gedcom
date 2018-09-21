package main

import (
	"fmt"
	"github.com/elliotchance/gedcom/html"
)

type navPillsRow struct {
	links []fmt.Stringer
}

func newNavPillsRow(links []fmt.Stringer) *navPillsRow {
	return &navPillsRow{
		links: links,
	}
}

func (c *navPillsRow) String() string {
	pills := newNavPills(c.links)
	div := html.NewDiv("", pills)
	column := html.NewColumn(html.EntireRow, div)
	row := html.NewRow(column)

	return row.String()
}
