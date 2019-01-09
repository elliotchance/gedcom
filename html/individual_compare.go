package html

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/util"
	"io"
)

type IndividualCompare struct {
	comparison  *gedcom.IndividualComparison
	filterFlags *util.FilterFlags
}

func NewIndividualCompare(comparison *gedcom.IndividualComparison, filterFlags *util.FilterFlags) *IndividualCompare {
	return &IndividualCompare{
		comparison:  comparison,
		filterFlags: filterFlags,
	}
}

func (c *IndividualCompare) appendChildren(nd *gedcom.NodeDiff, prefix string) []Component {
	title := prefix + nd.Tag().String()
	row := NewDiffRow(title, nd, c.filterFlags.HideEqual)
	tableRows := []Component{row}

	for _, child := range nd.Children {
		children := c.appendChildren(child, prefix+"&nbsp;&nbsp;&nbsp;&nbsp;")
		tableRows = append(tableRows, children...)
	}

	return tableRows
}

func (c *IndividualCompare) WriteTo(w io.Writer) (int64, error) {
	left := c.comparison.Left
	right := c.comparison.Right

	var name Component = nil

	if n := left; n != nil {
		name = NewIndividualNameAndDates(n, true, "")
	}

	if n := right; name == nil && n != nil {
		name = NewIndividualNameAndDates(n, true, "")
	}

	if name == nil {
		name = NewEmpty()
	}

	if !gedcom.IsNil(left) {
		left = c.filterFlags.Filter(left).(*gedcom.IndividualNode)
	}

	if !gedcom.IsNil(right) {
		right = c.filterFlags.Filter(right).(*gedcom.IndividualNode)
	}

	diff := gedcom.CompareNodes(left, right)

	diff.Sort()

	tableRows := c.appendChildren(diff, "")

	// Parents
	leftParents := gedcom.IndividualNodes{}
	if !gedcom.IsNil(left) {
		for _, parents := range left.Parents() {
			if parent := parents.Husband(); parent != nil {
				leftParents = append(leftParents, parent)
			}
			if parent := parents.Wife(); parent != nil {
				leftParents = append(leftParents, parent)
			}
		}
	}

	rightParents := gedcom.IndividualNodes{}
	if !gedcom.IsNil(right) {
		for _, parents := range right.Parents() {
			if parent := parents.Husband(); parent != nil {
				rightParents = append(rightParents, parent)
			}
			if parent := parents.Wife(); parent != nil {
				rightParents = append(rightParents, parent)
			}
		}
	}

	compareOptions := gedcom.NewIndividualNodesCompareOptions()
	for _, parents := range leftParents.Compare(rightParents, compareOptions) {
		var row *DiffRow
		name := "Parent"

		switch {
		case !gedcom.IsNil(parents.Left) && !gedcom.IsNil(parents.Right):
			row = NewDiffRow(name, &gedcom.NodeDiff{
				Left:  parents.Left.Name(),
				Right: parents.Right.Name(),
			}, c.filterFlags.HideEqual)

		case !gedcom.IsNil(parents.Left):
			row = NewDiffRow(name, &gedcom.NodeDiff{
				Left: parents.Left.Name(),
			}, c.filterFlags.HideEqual)

		case !gedcom.IsNil(parents.Right):
			row = NewDiffRow(name, &gedcom.NodeDiff{
				Right: parents.Right.Name(),
			}, c.filterFlags.HideEqual)
		}

		tableRows = append(tableRows, row)
	}

	// Spouses
	switch {
	case !gedcom.IsNil(left) && !gedcom.IsNil(right):
		for _, spouse := range left.Spouses().Compare(right.Spouses(), compareOptions) {
			nodeDiff := &gedcom.NodeDiff{}

			if spouse.Left != nil {
				nodeDiff.Left = spouse.Left.Name()
			}

			if spouse.Right != nil {
				nodeDiff.Right = spouse.Right.Name()
			}

			row := NewDiffRow("Spouse", nodeDiff, c.filterFlags.HideEqual)
			tableRows = append(tableRows, row)
		}

	case !gedcom.IsNil(left):
		for _, spouse := range left.Spouses() {
			row := NewDiffRow("Spouse", &gedcom.NodeDiff{
				Left: spouse.Name(),
			}, c.filterFlags.HideEqual)
			tableRows = append(tableRows, row)
		}

	case !gedcom.IsNil(right):
		for _, spouse := range right.Spouses() {
			row := NewDiffRow("Spouse", &gedcom.NodeDiff{
				Right: spouse.Name(),
			}, c.filterFlags.HideEqual)
			tableRows = append(tableRows, row)
		}
	}

	leftAnchor := ""
	rightAnchor := ""

	if c.comparison.Left != nil {
		leftAnchor = c.comparison.Left.Pointer()
	}

	if c.comparison.Right != nil {
		rightAnchor = c.comparison.Right.Pointer()
	}

	return NewComponents(
		NewAnchor(leftAnchor),
		NewAnchor(rightAnchor),
		NewBigTitle(1, name),
		NewSpace(),
		NewTable("", tableRows...),
	).WriteTo(w)
}
