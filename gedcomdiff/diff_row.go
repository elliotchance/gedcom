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
	if c.hideSame {
		if c.nd.IsDeepEqual() {
			return ""
		}

		if c.nd.Tag().IsEvent() && len(c.nd.Children) == 0 {
			return ""
		}
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
		html.NewTableCell(html.NewText(c.name)),
		html.NewTableCell(html.NewText(left)).Class(leftClass).Style("width: 40%"),
		html.NewTableCell(html.NewText(right)).Class(rightClass).Style("width: 40%"),
	).String()
}
