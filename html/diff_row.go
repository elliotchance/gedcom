package html

import (
	"github.com/elliotchance/gedcom"
	"io"
)

type DiffRow struct {
	name     string
	nd       *gedcom.NodeDiff
	hideSame bool
}

func NewDiffRow(name string, nd *gedcom.NodeDiff, hideSame bool) *DiffRow {
	return &DiffRow{
		name:     name,
		nd:       nd,
		hideSame: hideSame,
	}
}

func (c *DiffRow) WriteTo(w io.Writer) (int64, error) {
	if c.hideSame {
		if c.nd.IsDeepEqual() {
			return writeNothing()
		}

		if c.nd.Tag().IsEvent() && len(c.nd.Children) == 0 {
			return writeNothing()
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

	return NewTableRow(
		NewTableCell(NewText(c.name)),
		NewTableCell(NewText(left)).Class(leftClass).Style("width: 40%"),
		NewTableCell(NewText(right)).Class(rightClass).Style("width: 40%"),
	).WriteTo(w)
}
