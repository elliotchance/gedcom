package html

import (
	"github.com/elliotchance/gedcom"
	"io"
	"fmt"
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

func (c *DiffRow) isEmpty() bool {
	if c.hideSame {
		if c.nd.IsDeepEqual() {
			return true
		}

		if c.nd.Tag().IsEvent() && len(c.nd.Children) == 0 {
			return true
		}
	}

	return false
}

func (c *DiffRow) valueAndPointer(node gedcom.Node) string {
	v := node.Value()

	if i, ok := node.(*gedcom.IndividualNode); ok {
		v = i.Name().String()
	}

	if node.Pointer() != "" {
		v += fmt.Sprintf(" <%s>", node.Pointer())
	}

	return v
}

func (c *DiffRow) WriteTo(w io.Writer) (int64, error) {
	if c.isEmpty() {
		return writeNothing()
	}

	leftClass := ""
	rightClass := ""

	left := ""
	right := ""

	switch {
	case gedcom.IsNil(c.nd.Left) && gedcom.IsNil(c.nd.Right):
		// do nothing

	case gedcom.IsNil(c.nd.Left):
		right = c.valueAndPointer(c.nd.Right)
		rightClass = "bg-primary"

	case gedcom.IsNil(c.nd.Right):
		left = c.valueAndPointer(c.nd.Left)
		leftClass = "bg-warning"

	default:
		if !c.nd.IsDeepEqual() {
			leftClass = "bg-info"
			rightClass = "bg-info"
		}
		left = c.valueAndPointer(c.nd.Left)
		right = c.valueAndPointer(c.nd.Right)
	}

	return NewTableRow(
		NewTableCell(NewText(c.name)),
		NewTableCell(NewText(left)).Class(leftClass).Style("width: 40%"),
		NewTableCell(NewText(right)).Class(rightClass).Style("width: 40%"),
	).WriteTo(w)
}
