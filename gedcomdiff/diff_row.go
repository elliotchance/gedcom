package main

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
)

type diffRow struct {
	name     string
	nd       *gedcom.NodeDiff
	hideSame bool
}

func newDiffRow(name string, nd *gedcom.NodeDiff, hideSame bool) *diffRow {
	return &diffRow{
		name:     name,
		nd:       nd,
		hideSame: hideSame,
	}
}

func (c *diffRow) String() string {
	if c.hideSame && c.nd.IsDeepEqual() {
		return ""
	}

	leftClass := ""
	rightClass := ""

	left := ""
	right := ""

	switch {
	case gedcom.IsNil(c.nd.Left) && gedcom.IsNil(c.nd.Right):
		// do nothing

	case gedcom.IsNil(c.nd.Left):
		right = c.nd.Right.Value()
		rightClass = "bg-primary"

	case gedcom.IsNil(c.nd.Right):
		left = c.nd.Left.Value()
		leftClass = "bg-warning"

	default:
		if !c.nd.IsDeepEqual() {
			leftClass = "bg-info"
			rightClass = "bg-info"
		}
		left = c.nd.Left.Value()
		right = c.nd.Right.Value()
	}

	return html.NewTableRow(
		html.NewStyledTableCell("", "", html.NewText(c.name)),
		html.NewStyledTableCell("width: 40%", leftClass, html.NewText(left)),
		html.NewStyledTableCell("width: 40%", rightClass, html.NewText(right)),
	).String()
}
