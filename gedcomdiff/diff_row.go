package main

import (
	"github.com/elliotchance/gedcom/html"
)

type diffRow struct {
	left, right, name string
	hideSame          bool
}

func newDiffRow(name, left, right string, hideSame bool) *diffRow {
	return &diffRow{
		name:     name,
		left:     left,
		right:    right,
		hideSame: hideSame,
	}
}

func (c *diffRow) String() string {
	if c.left == c.right && c.hideSame {
		return ""
	}

	leftClass := ""
	rightClass := ""

	switch {
	case c.left == "" && c.right == "":
		return ""

	case c.left == "":
		rightClass = "bg-primary"

	case c.right == "":
		leftClass = "bg-warning"

	case c.left != c.right:
		leftClass = "bg-info"
		rightClass = "bg-info"
	}

	return html.NewTableRow(
		html.NewStyledTableCell("", "", html.NewText(c.name)),
		html.NewStyledTableCell("width: 40%", leftClass, html.NewText(c.left)),
		html.NewStyledTableCell("width: 40%", rightClass, html.NewText(c.right)),
	).String()
}
